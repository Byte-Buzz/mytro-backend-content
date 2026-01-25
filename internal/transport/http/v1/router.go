package v1

import (
	"mytro-backend-content/internal/app"
	"mytro-backend-content/internal/transport/http/v1/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, app *app.App) {
	router.POST("/upload", handler.CreateContentUploadHandler(app))
	router.GET("/upload/:id", handler.GetContentUploadStatusHandler(app))
}
