package handler

import (
	"go-dev-sample/internal/application/dto"
	"go-dev-sample/internal/domain/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserList(c echo.Context) error {
	// サービスの呼び出し
	users, err := h.service.GetUserList()
	if err != nil {
		log.Printf("failed to get user list: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// レスポンスの生成
	response := dto.UserList{}
	for _, user := range users {
		response = append(response, &dto.User{
			ID:    strconv.Itoa(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	// リクエストパラメータの取得
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter")
	}

	// サービスの呼び出し
	user, err := h.service.GetUserByID(userId)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// レスポンスの生成
	response := &dto.User{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	// リクエストパラメータの取得
	req := new(dto.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		log.Printf("failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter")
	}

	// バリデーション
	if err := req.Validate(nil); err != nil {
		log.Printf("failed to validate request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter")
	}

	// サービスの呼び出し
	user, err := h.service.CreateUser(*req.Name, *req.Email)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// レスポンスの生成
	response := &dto.User{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, response)
}
