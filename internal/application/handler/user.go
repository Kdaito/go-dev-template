package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Kdaito/go-dev-template/internal/application/request"
	"github.com/Kdaito/go-dev-template/internal/application/response"
	"github.com/Kdaito/go-dev-template/internal/domain/service"
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
		log.Printf("failed to get user list: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
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
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter")
	}

	// サービスの呼び出し
	user, err := h.service.GetUserByID(userId)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
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
		log.Printf("failed to bind request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter")
	}

	// バリデーション
	if err := c.Validate(req); err != nil {
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
	res := &response.UserResponse{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, res)
}
