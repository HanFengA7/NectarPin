package user

import (
	"net/http"
	"strings"

	userservice "nectarpin/api/services/user"
	"nectarpin/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *userservice.UserService
}

func NewUserController(service *userservice.UserService) *UserController {
	return &UserController{service: service}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,len=32"`
	Email    string `json:"email" binding:"required,email"`
	Nickname string `json:"nickname" binding:"omitempty,max=100"`
}

type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Status   int16  `json:"status"`
	Role     int16  `json:"role"`
}

type LoginResponse struct {
	User  UserResponse               `json:"user"`
	Token *userservice.TokenResponse `json:"token"`
}

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
		Password: req.Password,
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
		case userservice.ErrPasswordFormat:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "密码格式错误",
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

func isEmail(account string) bool {
	return strings.Contains(account, "@")
}

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
	account := strings.TrimSpace(req.Account)

	var loginType string
	if isEmail(account) {
		loginType = "邮箱"
	} else {
		loginType = "用户名"
	}

	utils.Logger.Infof("用户", "登录尝试: %s=%s, IP: %s", loginType, account, ip)

	userData, token, err := c.service.LoginByAccount(account, req.Password, ip)
	if err != nil {
		switch err {
		case userservice.ErrUserNotFound:
			utils.Logger.Infof("用户", "登录失败: 用户不存在 (%s=%s, IP: %s)", loginType, account, ip)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "用户不存在或密码错误",
			})
		case userservice.ErrInvalidPassword:
			utils.Logger.Infof("用户", "登录失败: 密码错误 (%s=%s, IP: %s)", loginType, account, ip)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "用户不存在或密码错误",
			})
		case userservice.ErrPasswordFormat:
			utils.Logger.Infof("用户", "登录失败: 密码格式错误 (%s=%s, IP: %s)", loginType, account, ip)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "密码格式错误",
			})
		case userservice.ErrUserDisabled:
			utils.Logger.Infof("用户", "登录失败: 用户已禁用 (%s=%s, IP: %s)", loginType, account, ip)
			ctx.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "用户已被禁用",
			})
		default:
			utils.Logger.Errorf("用户", "登录失败: %v (%s=%s, IP: %s)", err, loginType, account, ip)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "登录失败",
			})
		}
		return
	}

	utils.Logger.Infof("用户", "用户登录成功: %s (通过%s, IP: %s)", userData.Username, loginType, ip)

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

func (c *UserController) GetProfile(ctx *gin.Context) {
	// 从上下文中获取用户 ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	// 调用服务层获取用户信息
	user, err := c.service.GetProfile(userID.(uint64))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "用户不存在",
		})
		return
	}

	// 返回用户信息（排除敏感字段）
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Status:   user.Status,
			Role:     user.Role,
		},
	})
}

func (c *UserController) Logout(ctx *gin.Context) {
	// 从上下文中获取用户 ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	// 从请求头获取 Access Token
	accessToken := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	if accessToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "缺少访问令牌",
		})
		return
	}

	// 验证并解析 Token
	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的令牌",
		})
		return
	}

	// 验证 Token 中的用户 ID 与上下文一致
	if claims.UserID != userID.(uint64) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "令牌与用户不匹配",
		})
		return
	}

	// 调用服务层执行登出逻辑（可选：将 Token 加入黑名单）
	// 注意：由于 Access Token 是短期的（1 小时），通常不需要特别处理
	// 主要是清除客户端的 Refresh Token

	utils.Logger.Infof("用户", "用户登出成功 (user_id=%d)", claims.UserID)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登出成功",
	})
}

func (c *UserController) LogoutByRefreshToken(ctx *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	err := c.service.LogoutByRefreshToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "登出失败",
		})
		return
	}

	utils.Logger.Infof("用户", "用户通过 Refresh Token 登出成功")

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登出成功",
	})
}

func (c *UserController) RevokeAllSessions(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	err := c.service.RevokeAllUserSessions(userID.(uint64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "撤销会话失败",
		})
		return
	}

	utils.Logger.Infof("用户", "用户撤销了所有会话 (user_id=%d)", userID)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "已撤销所有其他设备的登录会话",
	})
}
