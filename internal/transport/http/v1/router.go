package v1

import (
	"mytro-backend-tile/internal/app"
	"mytro-backend-tile/internal/transport/http/v1/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, app *app.App) {
	router.POST("/uploads", handler.CreateUploadHandler(app))
	router.GET("/uploads", handler.GetUploadListHandler(app))
	router.GET("/uploads/:id", handler.GetUploadStatusHandler(app))
}
