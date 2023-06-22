package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NoCachedMiddleware(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func CorsMiddleware(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, orgin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
		c.AbortWithStatus(200)
	}
}
