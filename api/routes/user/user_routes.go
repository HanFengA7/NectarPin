package user

import (
	usercontroller "nectarpin/api/controllers/user"
	"nectarpin/api/middlewares"
	userrepo "nectarpin/api/repositories/user"
	userservice "nectarpin/api/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	protectedAPI := router.Group("/api/protected")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		userV1 := protectedAPI.Group("/user/v1")
		{
			userV1.GET("/profile", userController.GetProfile)
		}
	}
}
