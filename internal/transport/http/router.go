package http

import (
	"mytro-backend-content/internal/app"
	"mytro-backend-content/internal/transport/http/middleware"
	v1 "mytro-backend-content/internal/transport/http/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(app *app.App) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Load public key for authentication
	publicKey, err := app.GetPublicKey()
	if err != nil {
		app.Logger.Panic("parse public key error", zap.Error(err))
	}

	// Apply middlewares
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger(app.Logger))
	router.Use(middleware.CORS(app.Config.CORS))
	router.Use(middleware.SecurityHeaders())
	router.Use(gin.Recovery())
	router.Use(middleware.Auth(publicKey, app.Logger))

	// Register health check endpoint
	registerHealth(router, app)

	baseGroup := router.Group("/api/content")

	// Register API v1 routes
	v1Group := baseGroup.Group("/v1")
	v1.RegisterRoutes(v1Group, app)

	// Register base API routes
	v1.RegisterRoutes(baseGroup, app)

	return router
}
