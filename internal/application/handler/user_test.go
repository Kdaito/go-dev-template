package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Kdaito/kinodokuna-be/internal/domain/model"
	"github.com/Kdaito/kinodokuna-be/internal/lib"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// IUserServiceのモック
type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) GetUserList() ([]*model.User, error) {
	ret := m.Called()
	return ret.Get(0).([]*model.User), ret.Error(1)
}
func (m *mockUserService) GetUserByID(id int) (*model.User, error) {
	ret := m.Called(id)
	user, ok := ret.Get(0).(*model.User)
	if !ok {
		return nil, ret.Error(1)
	}
	return user, ret.Error(1)
}
func (m *mockUserService) CreateUser(name, email string) (*model.User, error) {
	ret := m.Called(name, email)
	user, ok := ret.Get(0).(*model.User)
	if !ok {
		return nil, ret.Error(1)
	}
	return user, ret.Error(1)
}

func TestGetUserList(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)
		mockedUsers := []*model.User{
			{
				ID:    1,
				Name:  "John Doe",
				Email: "John@example.com",
			},
			{
				ID:    2,
				Name:  "Kathy Smith",
				Email: "Kathy@example.com",
			},
		}
		service.On("GetUserList").Return(mockedUsers, nil)

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		var expectedUserJson = `[
			{
				"id": "1",
				"name": "John Doe",
				"email": "John@example.com"
			},
			{
				"id": "2",
				"name": "Kathy Smith",
				"email": "Kathy@example.com"
			}
		]`

		// assertion
		if assert.NoError(t, userHandler.GetUserList(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, expectedUserJson, rec.Body.String())
		}
	})

	tt.Run("異常系_サービスでエラー", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)
		mockedUsers := []*model.User{}
		service.On("GetUserList").Return(mockedUsers, assert.AnError)

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		err := userHandler.GetUserList(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})
}

func TestGetUserByID(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		targetUserId := 1

		// service mock setup
		service := new(mockUserService)
		mockedUser := &model.User{
			ID:    targetUserId,
			Name:  "John Doe",
			Email: "John@example.com",
		}
		service.On("GetUserByID", targetUserId).Return(mockedUser, nil)

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(targetUserId))

		userHandler := NewUserHandler(service)

		var expectedUserJson = `{
			"id": "1",
			"name": "John Doe",
			"email": "John@example.com"
		}`

		// assertion
		if assert.NoError(t, userHandler.GetUserByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, expectedUserJson, rec.Body.String())
		}
	})

	tt.Run("異常系_リクエストパラメータが不正", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// リクエストパラメータが不正なので、idに文字列を指定
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("invalid")

		userHandler := NewUserHandler(service)

		err := userHandler.GetUserByID(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	tt.Run("異常系_サービスでエラー", func(t *testing.T) {
		targetUserId := 1

		// service mock setup
		service := new(mockUserService)
		service.On("GetUserByID", targetUserId).Return(nil, assert.AnError)

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(targetUserId))

		userHandler := NewUserHandler(service)

		err := userHandler.GetUserByID(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})
}

func TestCreateUser(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)
		mockedUser := &model.User{
			ID:    1,
			Name:  "Kathy Smith",
			Email: "Kathy@example.com",
		}
		service.On("CreateUser", mockedUser.Name, mockedUser.Email).Return(mockedUser, nil)

		var requestBody = `{
			"name": "Kathy Smith",
			"email": "Kathy@example.com"
		}`

		// request setup
		e := echo.New()
		e.Validator = &lib.CustomValidator{Validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		var expectedUserJson = `{
				"id": "1",
				"name": "Kathy Smith",
				"email": "Kathy@example.com"
			}`

		// assertion
		if assert.NoError(t, userHandler.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.JSONEq(t, expectedUserJson, rec.Body.String())
		}
	})

	tt.Run("異常系_リクエストパラメータが不正", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)

		// 不正なJSONリクエスト
		var requestBody = `{
			"name": "Kathy Smith"
			"email": "Kathy@example.com"
		}`

		// request setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		err := userHandler.CreateUser(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	tt.Run("異常系_リクエストバリデーションで失敗", func(t *testing.T) {
		// service mock setup
		service := new(mockUserService)

		// バリデーションに失敗するJSONリクエスト
		var requestBody = `{
			"name": "",
			"email": "invalid email"
		}`

		// request setup
		e := echo.New()
		e.Validator = &lib.CustomValidator{Validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		err := userHandler.CreateUser(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	tt.Run("異常系_サービスでエラー", func(t *testing.T) {
		var newUserName = "Kathy Smith"
		var newUserEmail = "Kathy@example.com"

		// service mock setup
		service := new(mockUserService)
		service.On("CreateUser", newUserName, newUserEmail).Return(nil, assert.AnError)

		var requestBody = fmt.Sprintf(`{
			"name": "%s",
			"email": "%s"
		}`, newUserName, newUserEmail)

		// request setup
		e := echo.New()
		e.Validator = &lib.CustomValidator{Validator: validator.New()}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserHandler(service)

		err := userHandler.CreateUser(c)

		// assertion
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})
}
