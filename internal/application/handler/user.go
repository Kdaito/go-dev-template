package handler

import (
	"net/http"
	"strconv"

	"github.com/Kdaito/kinodokuna-be/internal/application/request"
	"github.com/Kdaito/kinodokuna-be/internal/application/response"
	"github.com/Kdaito/kinodokuna-be/internal/domain/service"
	"github.com/Kdaito/kinodokuna-be/internal/lib/errors"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserList(c echo.Context) error {
	// サービスの呼び出し
	users, err := h.service.GetUserList()
	if err != nil {
		return err
	}

	// レスポンスの生成
	res := response.UserList{}
	for _, user := range users {
		res = append(res, &response.UserResponse{
			ID:    strconv.Itoa(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	// リクエストパラメータの取得
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.New(http.StatusBadRequest, err.Error())
	}

	// サービスの呼び出し
	user, err := h.service.GetUserByID(userId)
	if err != nil {
		return err
	}

	// レスポンスの生成
	res := &response.UserResponse{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	// リクエストパラメータの取得
	req := new(request.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return errors.New(http.StatusBadRequest, err.Error())
	}

	// バリデーション
	if err := c.Validate(req); err != nil {
		return errors.New(http.StatusBadRequest, err.Error())
	}

	// サービスの呼び出し
	user, err := h.service.CreateUser(*req.Name, *req.Email)
	if err != nil {
		return err
	}

	// レスポンスの生成
	res := &response.UserResponse{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, res)
}
