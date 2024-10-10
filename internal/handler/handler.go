package handler

import (
	"bank-app-backend/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	API_V1_ROUTE = "/api/v1"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group(API_V1_ROUTE)

	apiV1.POST("/signUp", h.SignUp)

	return router
}
