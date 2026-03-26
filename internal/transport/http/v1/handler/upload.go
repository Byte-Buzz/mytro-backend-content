package handler

import (
	"mytro-backend-content/internal/app"
	"mytro-backend-content/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func ConfirmContentUploadHandler(app *app.App) gin.HandlerFunc {
	services := app.Services
	return func(c *gin.Context) {
		contentID, err := uuid.Parse(c.Param("content_id"))
		if err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			return
		}

		userId := middleware.AuthFromContext(c)
		if userId == uuid.Nil {
			c.JSON(401, gin.H{"status": "error", "message": "unauthorized"})
		}

		valid, err := services.ContentService.CheckContentOwner(c, contentID, userId)
		if err != nil {
			c.JSON(403, gin.H{"status": "error", "message": err.Error()})
			return
		}

		if !valid {
			c.JSON(403, gin.H{"status": "error", "message": "forbidden"})
		}

		err = services.ContentService.CompleteUpload(c, contentID)
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	}
}

func GetContentUploadStatusHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
