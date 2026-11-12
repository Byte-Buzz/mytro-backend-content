package handler

import (
	"mytro-backend-tile/internal/app"
	"mytro-backend-tile/internal/app/service"
	"mytro-backend-tile/internal/domain/models"
	"mytro-backend-tile/internal/transport/http/middleware"
	"mytro-backend-tile/internal/transport/http/v1/request"
	"mytro-backend-tile/internal/transport/http/v1/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateContentUploadHandler(app *app.App) gin.HandlerFunc {
	services := app.Services

	return func(c *gin.Context) {
		var createRequest request.CreateUploadRequest

		err := c.BindJSON(&createRequest)
		if err != nil {
			c.JSON(400, gin.H{"status": "error", "message": err.Error()})
			return
		}

		userID := middleware.AuthFromContext(c)
		if userID == uuid.Nil {
			c.JSON(401, gin.H{"status": "error", "message": "unauthorized"})
			return
		}

		url, err := services.ContentService.CreateUpload(c.Request.Context(), models.CreateUpload{
			FileType: createRequest.ContentType,
			OwnerID:  userID,
		})
		if err == service.ErrUnsupportedContentType {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "error", "message": "unsupported file type"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}

		body := response.CreateUploadResponse{
			ContentID: url.ContentID.String(),
			UploadURL: url.UploadURL,
			ExpiresAt: url.ExpiresAt,
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": body})
	}
}

func GetContentUploadStatusHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"status": "error", "message": "not implemented"})
	}
}

func GetContentUploadListHandler(app *app.App) gin.HandlerFunc {
	services := app.Services

	return func(c *gin.Context) {
		userID := middleware.AuthFromContext(c)
		if userID == uuid.Nil {
			c.JSON(401, gin.H{"status": "error", "message": "unauthorized"})
			return
		}

		contentList, err := services.ContentService.GetNewContentList(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": contentList})
	}
}
