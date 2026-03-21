package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var allowedOrigins = map[string]struct{}{
	"http://localhost:5173": {},
	"http://127.0.0.1:5173": {},
	"http://localhost:4173": {},
	"http://127.0.0.1:4173": {},
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		if _, allowed := allowedOrigins[origin]; allowed {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Vary", "Origin")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		}

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
