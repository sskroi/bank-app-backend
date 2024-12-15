package apihandler

import (
	"bank-app-backend/internal/service"
	"errors"
	"os"

	"github.com/gin-contrib/cors"
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

// TODO: return detailed error messages from API

func (h *Handler) InitRoutes() *gin.Engine {
	appMode := os.Getenv("BANK_APP_MODE")
	if appMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		// hard coded
		// AllowOrigins:     []string{"https://iorkss.ru"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{},
		// AllowCredentials: true,
	}))

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		// only authorized routes
		authOnly := v1.Group("", h.verifyAuth)
		authOnly.POST("/auth/check", h.check)

		user := authOnly.Group("/user")
		{
			user.POST("/update-profile", h.updateUserProfile)
		}

		account := authOnly.Group("/account")
		{
			account.POST("", h.createAccount)
			account.DELETE("", h.closeAccount)
		}
		accounts := authOnly.Group("/accounts")
		{
			accounts.GET("", h.userAccounts)
		}

		transaction := authOnly.Group("/transaction")
		{
			transaction.POST("", h.createTransaction)
		}

		// SWAGGER
		appMode := os.Getenv("BANK_APP_SWAGGER_ENABLED")
		if appMode == "1" {
			// available on {HOST}/api/v1/swagger/index.html
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		}
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
