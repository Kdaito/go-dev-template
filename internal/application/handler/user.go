package handler

import (
	"go-dev-sample/internal/application/response"
	"go-dev-sample/internal/domain/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	// リクエストパラメータの取得
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}

	// サービスの呼び出し
	user, err := h.service.GetUserByID(userId)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// レスポンスの生成
	response := response.NewGetUserByIdResponse(user)
	c.JSON(http.StatusOK, response)
}
