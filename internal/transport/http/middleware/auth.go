package middleware

import (
	"context"
	"crypto/rsa"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type authKeyType struct{}

var authIDKey = authKeyType{}

func Auth(publicKey *rsa.PublicKey, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				logger.Error(
					"unexpected signing method",
					zap.String("request_id", RequestIDFromContext(c)),
					zap.String("method", token.Header["alg"].(string)),
				)
				return nil, jwt.ErrTokenMalformed
			}
			return publicKey, nil
		})

		if err != nil {
			logger.Warn(
				"invalid token",
				zap.String("request_id", RequestIDFromContext(c)),
				zap.Error(err),
			)
			c.Next()
			return
		}

		if !token.Valid {
			c.Next()
			return
		}

		id, err := token.Claims.GetSubject()
		if err != nil {
			logger.Error("parse claims error", zap.Error(err), zap.String("request_id", RequestIDFromContext(c)))
			c.Next()
			return
		}

		if err := uuid.Validate(id); err != nil {
			logger.Error("invalid user id", zap.Error(err), zap.String("request_id", RequestIDFromContext(c)))
			c.Next()
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(
			c.Request.Context(),
			authIDKey,
			id,
		))

		c.Next()
	}
}

func AuthFromContext(ctx context.Context) uuid.UUID {
	id, ok := ctx.Value(authIDKey).(string)

	if !ok {
		return uuid.Nil
	}

	userId, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil
	}

	return userId
}
