// Package user 提供用户路由配置
// 定义用户相关的 HTTP 路由和路由组
package user

import (
	usercontroller "nectarpin/api/controllers/user"
	userrepo "nectarpin/api/repositories/user"
	userservice "nectarpin/api/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 配置用户相关路由
// 参数:
//   - router: Gin 路由引擎
//   - db: 数据库连接实例
//
// 路由结构:
//
//	/api/public/user/v1
//	  ├── POST /register  用户注册
//	  └── POST /login     用户登录
//
// 该函数会初始化完整的依赖链:
//  1. 创建 UserRepository (数据访问层)
//  2. 创建 UserService (业务逻辑层)
//  3. 创建 UserController (控制器层)
//  4. 注册路由
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewUserController(userService)

	publicAPI := router.Group("/api/public")
	{
		userV1 := publicAPI.Group("/user/v1")
		{
			userV1.POST("/register", userController.Register)
			userV1.POST("/login", userController.Login)
		}
	}
}
