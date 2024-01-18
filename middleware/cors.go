package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "*"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}
