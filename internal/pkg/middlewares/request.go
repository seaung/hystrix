package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/seaung/hystrix/internal/pkg/known"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.RequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(known.RequestIDKey, requestID)
		c.Writer.Header().Set(known.RequestIDKey, requestID)
		c.Next()
	}
}
