// Package user 提供用户控制器层实现
// 处理 HTTP 请求，调用服务层，返回响应
package user

import (
	"net/http"

	userservice "nectarpin/api/services/user"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
// 处理用户相关的 HTTP 请求
type UserController struct {
	service *userservice.UserService
}

// NewUserController 创建用户控制器实例
// 参数:
//   - service: 用户服务实例
//
// 返回:
//   - *UserController: 用户控制器实例
func NewUserController(service *userservice.UserService) *UserController {
	return &UserController{service: service}
}

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"` // 用户名，必填，3-50字符
	Email    string `json:"email" binding:"required,email"`           // 邮箱，必填，邮箱格式
	Nickname string `json:"nickname" binding:"omitempty,max=100"`     // 昵称，可选，最大100字符
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名，必填
}

// UserResponse 用户响应结构体
type UserResponse struct {
	ID       uint64 `json:"id"`       // 用户 ID
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像 URL
	Status   int16  `json:"status"`   // 账户状态
	Role     int16  `json:"role"`     // 用户角色
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	User  UserResponse               `json:"user"`  // 用户信息
	Token *userservice.TokenResponse `json:"token"` // 令牌信息
}

// Register 处理用户注册请求
// 路由: POST /api/public/user/v1/register
//
// 请求体:
//
//	{
//	    "username": "用户名",
//	    "email": "邮箱地址",
//	    "nickname": "昵称(可选)"
//	}
//
// 成功响应: 201 Created
//
//	{
//	    "code": 201,
//	    "message": "注册成功",
//	    "data": { ... }
//	}
func (c *UserController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	input := &userservice.RegisterInput{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
	}

	userData, err := c.service.Register(input)
	if err != nil {
		switch err {
		case userservice.ErrUsernameExists:
			ctx.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "用户名已存在",
			})
		case userservice.ErrEmailExists:
			ctx.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "邮箱已存在",
			})
		default:
			utils.Logger.Errorf("用户", "注册失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "注册失败",
			})
		}
		return
	}

	utils.Logger.Infof("用户", "用户注册成功: %s", userData.Username)

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "注册成功",
		"data": UserResponse{
			ID:       userData.ID,
			Username: userData.Username,
			Email:    userData.Email,
			Nickname: userData.Nickname,
			Avatar:   userData.Avatar,
			Status:   userData.Status,
			Role:     userData.Role,
		},
	})
}

// Login 处理用户登录请求
// 路由: POST /api/public/user/v1/login
//
// 请求体:
//
//	{
//	    "username": "用户名"
//	}
//
// 成功响应: 200 OK
//
//	{
//	    "code": 200,
//	    "message": "登录成功",
//	    "data": { ... }
//	}
func (c *UserController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	ip := ctx.ClientIP()

	userData, token, err := c.service.Login(req.Username, ip)
	if err != nil {
		switch err {
		case userservice.ErrUserNotFound:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "用户不存在",
			})
		case userservice.ErrUserDisabled:
			ctx.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "用户已被禁用",
			})
		default:
			utils.Logger.Errorf("用户", "登录失败: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "登录失败",
			})
		}
		return
	}

	utils.Logger.Infof("用户", "用户登录成功: %s (IP: %s)", userData.Username, ip)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": LoginResponse{
			User: UserResponse{
				ID:       userData.ID,
				Username: userData.Username,
				Email:    userData.Email,
				Nickname: userData.Nickname,
				Avatar:   userData.Avatar,
				Status:   userData.Status,
				Role:     userData.Role,
			},
			Token: token,
		},
	})
}
