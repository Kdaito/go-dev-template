package handler

import (
	"go-dev-sample/internal/application/dto"
	"go-dev-sample/internal/domain/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
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
