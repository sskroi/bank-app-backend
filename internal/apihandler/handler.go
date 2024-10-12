package apihandler

import (
	"bank-app-backend/internal/service"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "bank-app-backend/docs"

	// gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
)

const (
	// this key sets by verifyAuth middleware, has type uuid.UUID
	userPubIdKey = "publicId"
)

type Handler struct {
	service *service.Services
}

func New(service *service.Services) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/sign-up", h.signUp)
		v1.POST("/auth/sign-in", h.signIn)

		// only authorized routes
		auth := v1.Group("/", h.verifyAuth)

		auth.POST("/user/update-profile", h.updateUserProfile)

		// SWAGGER
		// available on localhost:8080/api/v1/swagger/index.html
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}

// getUserPubId extracs `userPubId` key set in auth middleware from gin.Context 
func getUserPubId(c *gin.Context) (uuid.UUID, error) {
	userPubIdAny, ok := c.Get(userPubIdKey)
	if !ok {
		return uuid.UUID{}, errors.New("can't get userPubId key from gin.Context")
	}

	userPubId, ok := userPubIdAny.(uuid.UUID)
	if !ok {
		return uuid.UUID{}, errors.New("can't cast userPubId to uuid.UUID")
	}

	return userPubId, nil
}
