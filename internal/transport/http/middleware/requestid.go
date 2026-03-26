package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestIDKeyType struct{}

var requestIDKey = requestIDKeyType{}

// RequestID is a middleware that assigns a unique request ID to each incoming HTTP request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewString()
		c.Header("X-Request-ID", uuid)
		c.Set("request_id", uuid)

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), requestIDKey, uuid))

		c.Next()
	}
}

func RequestIDFromContext(ctx context.Context) string {
	id := ctx.Value(requestIDKey).(string)
	return id
}
