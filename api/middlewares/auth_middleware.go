package middlewares

import (
	"net/http"
	"strings"

	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未提供认证令牌",
			})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "认证令牌格式错误",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			if err == utils.ErrExpiredToken {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"code":    401,
					"message": "令牌已过期",
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"code":    401,
					"message": "无效的令牌",
				})
			}
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}

func GetUserID(ctx *gin.Context) (uint64, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(uint64), true
}
