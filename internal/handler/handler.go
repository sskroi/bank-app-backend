package handler

import (
	"bank-app-backend/internal/service"

	"github.com/gin-gonic/gin"

	_ "bank-app-backend/docs"

	// gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
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

	apiV1.POST("/auth/sign-up", h.SignUp)

	// SWAGGER
	// available on localhost:8080/api/v1/swagger/index.html
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
