package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/hystrix/internal/pkg/core"
	"github.com/seaung/hystrix/internal/pkg/errno"
	"github.com/seaung/hystrix/internal/pkg/known"
	"github.com/seaung/hystrix/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.TokenInvalidError, nil)
			c.Abort()
			return
		}
		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
