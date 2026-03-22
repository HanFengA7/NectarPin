package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 按允许的 Origin 白名单设置 CORS 响应头；列表在应用启动时固定。
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, o := range allowedOrigins {
		allowed[o] = struct{}{}
	}

	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		if _, ok := allowed[origin]; ok {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Vary", "Origin")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Refresh-Token")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		}

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
