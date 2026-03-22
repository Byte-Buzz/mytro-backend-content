package http

import (
	"mytro-backend-content/internal/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// registerHealth registers the health check endpoint to the router.
func registerHealth(router *gin.Engine, app *app.App) {
	router.GET("/health", func(c *gin.Context) {
		app.Logger.Info("Health check received")
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/ready", func(c *gin.Context) {
		app.Logger.Info("Readiness check received")
		sqlDB, err := app.DB.DB()
		if err != nil {
			app.Logger.Error("Failed to get database connection", zap.Error(err))
			c.JSON(500, gin.H{"status": "error", "message": "Database connection error"})
		}
		if err := sqlDB.Ping(); err != nil {
			app.Logger.Error("Database ping failed", zap.Error(err))
			c.JSON(500, gin.H{"status": "error", "message": "Database not reachable"})
			return
		}
		app.Logger.Info("Database connection is healthy")
		c.JSON(200, gin.H{"status": "ok"})
	})
}
