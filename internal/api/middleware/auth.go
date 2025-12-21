package middleware

import (
	"nectarpin/internal/config"
	"nectarpin/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT 认证中间件
func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头获取 token
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "请求头中缺少 Authorization",
			})
			ctx.Abort()
			return
		}

		// 验证 token 格式: Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "Authorization 格式错误，正确格式: Bearer <token>",
			})
			ctx.Abort()
			return
		}

		// 解析 token
		claims, err := utils.ParseToken(parts[1], cfg.Server.JWTSecret)
		if err != nil {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "无效的 token: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		// 将用户信息存储到上下文中
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_name", claims.UserName)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

// GetUserID 从上下文中获取用户ID
func GetUserID(ctx *gin.Context) (uint, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(uint), true
}

// GetUserName 从上下文中获取用户名
func GetUserName(ctx *gin.Context) (string, bool) {
	userName, exists := ctx.Get("user_name")
	if !exists {
		return "", false
	}
	return userName.(string), true
}

// GetRole 从上下文中获取用户角色
func GetRole(ctx *gin.Context) (string, bool) {
	role, exists := ctx.Get("role")
	if !exists {
		return "", false
	}
	return role.(string), true
}

// RequireRole 角色验证中间件
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := GetRole(ctx)
		if !exists {
			ctx.JSON(403, gin.H{
				"code": 403,
				"msg":  "无法获取用户角色",
			})
			ctx.Abort()
			return
		}

		// 检查用户角色是否在允许的角色列表中
		hasPermission := false
		for _, role := range roles {
			if userRole == role {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			ctx.JSON(403, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
