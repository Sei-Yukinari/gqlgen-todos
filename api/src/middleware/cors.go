package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		})
		c.Next()
	}
}

func applyCors(c *gin.Context, allowOrigins []string) {
	origin := c.Request.Header.Get("Origin")
	for _, value := range allowOrigins {
		if value == origin {
			c.Header("Access-Control-Allow-Origin", origin)
		}
	}
}
