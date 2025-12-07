package user

import (
	"nectarpin/internal/database"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	// 使用全局数据库实例
	db := database.GetDB()
	// 或者直接使用 database.DB
	// db := database.DB
	
	// 示例：创建用户
	// db.Create(&user)
	_ = db // 避免未使用变量错误
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	// 使用全局数据库实例
	db := database.GetDB()
	
	// 示例：查询用户
	// db.Where("username = ?", username).First(&user)
	_ = db // 避免未使用变量错误
}
