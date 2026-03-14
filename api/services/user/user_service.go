// Package user 提供用户业务逻辑层实现
// 封装用户注册、登录、令牌管理等核心业务逻辑
package user

import (
	"errors"
	"time"

	"nectarpin/api/models"
	userrepo "nectarpin/api/repositories/user"
	"nectarpin/internal/utils"
)

// UserService 用户服务
// 提供用户相关的业务逻辑处理
type UserService struct {
	repo *userrepo.UserRepository
}

// NewUserService 创建用户服务实例
// 参数:
//   - repo: 用户数据仓库实例
//
// 返回:
//   - *UserService: 用户服务实例
func NewUserService(repo *userrepo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// 业务错误定义
var (
	ErrUsernameExists = errors.New("用户名已存在") // 用户名重复错误
	ErrEmailExists    = errors.New("邮箱已存在")  // 邮箱重复错误
	ErrUserNotFound   = errors.New("用户不存在")  // 用户不存在错误
	ErrUserDisabled   = errors.New("用户已被禁用") // 用户禁用错误
	ErrInvalidToken   = errors.New("无效的令牌")  // 令牌无效错误
	ErrTokenExpired   = errors.New("令牌已过期")  // 令牌过期错误
)

// RegisterInput 注册请求输入参数
type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"` // 用户名，必填，3-50字符
	Email    string `json:"email" binding:"required,email"`           // 邮箱，必填，邮箱格式
	Nickname string `json:"nickname" binding:"omitempty,max=100"`     // 昵称，可选，最大100字符
}

// LoginInput 登录请求输入参数
type LoginInput struct {
	Username string `json:"username" binding:"required"` // 用户名，必填
}

// TokenResponse 令牌响应数据
type TokenResponse struct {
	AccessToken  string    `json:"access_token"`  // 访问令牌
	RefreshToken string    `json:"refresh_token"` // 刷新令牌
	ExpiresAt    time.Time `json:"expires_at"`    // 过期时间
	TokenType    string    `json:"token_type"`    // 令牌类型
}

// Register 用户注册
// 参数:
//   - input: 注册输入参数
//
// 返回:
//   - *models.User: 注册成功的用户数据
//   - error: 注册过程中的错误信息
//
// 业务逻辑:
//  1. 检查用户名是否已存在
//  2. 检查邮箱是否已存在
//  3. 创建用户记录
func (s *UserService) Register(input *RegisterInput) (*models.User, error) {
	exists, err := s.repo.ExistsByUsername(input.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUsernameExists
	}

	exists, err = s.repo.ExistsByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	nickname := input.Nickname
	if nickname == "" {
		nickname = input.Username
	}

	userData := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Nickname: nickname,
		Status:   models.UserStatusActive,
		Role:     models.UserRoleUser,
	}

	if err := s.repo.Create(userData); err != nil {
		return nil, err
	}

	return userData, nil
}

// Login 用户登录
// 参数:
//   - username: 用户名
//   - ip: 客户端 IP 地址
//
// 返回:
//   - *models.User: 登录成功的用户数据
//   - *TokenResponse: 生成的令牌信息
//   - error: 登录过程中的错误信息
//
// 业务逻辑:
//  1. 查询用户信息
//  2. 检查用户状态
//  3. 更新登录信息
//  4. 生成访问令牌
func (s *UserService) Login(username string, ip string) (*models.User, *TokenResponse, error) {
	userData, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, nil, ErrUserNotFound
	}

	if userData.Status == models.UserStatusDisabled {
		return nil, nil, ErrUserDisabled
	}

	if err := s.repo.UpdateLoginInfo(userData.ID, ip); err != nil {
		return nil, nil, err
	}

	tokenResponse, err := s.generateTokens(userData.ID)
	if err != nil {
		return nil, nil, err
	}

	return userData, tokenResponse, nil
}

// generateTokens 生成用户令牌
// 参数:
//   - userID: 用户唯一标识
//
// 返回:
//   - *TokenResponse: 生成的令牌信息
//   - error: 生成过程中的错误信息
//
// 生成访问令牌和刷新令牌，有效期分别为 24 小时和 7 天
func (s *UserService) generateTokens(userID uint64) (*TokenResponse, error) {
	accessToken, accessExpiresAt, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	accessTokenRecord := &models.UserToken{
		UserID:    userID,
		Token:     accessToken,
		Type:      models.TokenTypeAccess,
		ExpiresAt: accessExpiresAt,
	}

	refreshTokenRecord := &models.UserToken{
		UserID:    userID,
		Token:     refreshToken,
		Type:      models.TokenTypeRefresh,
		ExpiresAt: refreshExpiresAt,
	}

	if err := s.repo.CreateToken(accessTokenRecord); err != nil {
		return nil, err
	}

	if err := s.repo.CreateToken(refreshTokenRecord); err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExpiresAt,
		TokenType:    "Bearer",
	}, nil
}

// RefreshToken 刷新访问令牌
// 参数:
//   - refreshToken: 刷新令牌
//
// 返回:
//   - *TokenResponse: 新生成的令牌信息
//   - error: 刷新过程中的错误信息
//
// 业务逻辑:
//  1. 验证刷新令牌有效性
//  2. 撤销原刷新令牌
//  3. 生成新的访问令牌和刷新令牌
func (s *UserService) RefreshToken(refreshToken string) (*TokenResponse, error) {
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		if err == utils.ErrExpiredToken {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	token, err := s.repo.FindTokenByToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if token.Type != models.TokenTypeRefresh {
		return nil, ErrInvalidToken
	}

	if err := s.repo.RevokeToken(refreshToken); err != nil {
		return nil, err
	}

	return s.generateTokens(claims.UserID)
}

func (s *UserService) Logout(accessToken string) error {
	return s.repo.RevokeToken(accessToken)
}
