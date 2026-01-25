package handler

import (
	"mytro-backend-content/internal/app"

	"github.com/gin-gonic/gin"
)

func CreateContentUploadHandler(app *app.App) gin.HandlerFunc {
	services := app.Services

	return func(c *gin.Context) {
		url, err := services.ContentService.CreateUpload(c)
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "url": url.UploadURL})
	}
}

func GetContentUploadStatusHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
