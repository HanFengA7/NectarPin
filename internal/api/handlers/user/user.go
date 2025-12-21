package user

import (
	"nectarpin/internal/config"
	"nectarpin/internal/database"
	"nectarpin/internal/models"
	"nectarpin/internal/repository"
	"nectarpin/internal/service"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService *service.UserService
	config      *config.Config
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(cfg *config.Config) *UserHandler {
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	return &UserHandler{
		userService: userService,
		config:      cfg,
	}
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	NickName string `json:"nick_name"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRegister 用户注册
func UserRegister(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler := NewUserHandler(cfg)

		var req RegisterRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"code": 400,
				"msg":  "请求参数错误: " + err.Error(),
			})
			return
		}

		// 创建用户对象
		user := &models.User{
			UserName: req.UserName,
			Password: req.Password,
			Email:    req.Email,
			NickName: req.NickName,
			Role:     "0", // 默认普通用户
		}

		// 调用 Service 层创建用户
		msg, code := handler.userService.CreateUser(user)
		if code != 200 {
			ctx.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
			"data": gin.H{
				"user_id":   user.ID,
				"user_name": user.UserName,
				"email":     user.Email,
			},
		})
	}
}

// UserLogin 用户登录
func UserLogin(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler := NewUserHandler(cfg)

		var req LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"code": 400,
				"msg":  "请求参数错误: " + err.Error(),
			})
			return
		}

		// 调用 Service 层获取用户信息
		user, err := handler.userService.GetUserByUsername(req.UserName)
		if err != nil {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "用户名或密码错误",
			})
			return
		}

		// 验证密码
		msg, code := utils.CheckUserPassword(req.Password, user.Password)
		if code != 200 {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "用户名或密码错误",
			})
			return
		}

		// 生成 JWT Token
		token, err := utils.GenerateToken(
			user.ID,
			user.UserName,
			user.Role,
			cfg.Server.JWTSecret,
			cfg.Server.JWTExpire,
		)
		if err != nil {
			ctx.JSON(500, gin.H{
				"code": 500,
				"msg":  "生成 token 失败: " + err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
			"data": gin.H{
				"token":     token,
				"user_id":   user.ID,
				"user_name": user.UserName,
				"nick_name": user.NickName,
				"email":     user.Email,
				"role":      user.Role,
			},
		})
	}
}

// GetCurrentUser 获取当前登录用户信息
func GetCurrentUser(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler := NewUserHandler(cfg)

		// 从中间件中获取用户ID
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "未获取到用户信息",
			})
			return
		}

		// 获取用户详细信息
		user, err := handler.userService.GetUserByID(userID.(uint))
		if err != nil {
			ctx.JSON(404, gin.H{
				"code": 404,
				"msg":  "用户不存在",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": gin.H{
				"user_id":   user.ID,
				"user_name": user.UserName,
				"nick_name": user.NickName,
				"email":     user.Email,
				"role":      user.Role,
			},
		})
	}
}
