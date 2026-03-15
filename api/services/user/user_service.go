package user

import (
	"errors"
	"strings"
	"time"

	"nectarpin/api/models"
	userrepo "nectarpin/api/repositories/user"
	"nectarpin/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *userrepo.UserRepository
}

func NewUserService(repo *userrepo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var (
	ErrUsernameExists  = errors.New("用户名已存在")
	ErrEmailExists     = errors.New("邮箱已存在")
	ErrUserNotFound    = errors.New("用户不存在")
	ErrUserDisabled    = errors.New("用户已被禁用")
	ErrInvalidPassword = errors.New("密码错误")
	ErrInvalidToken    = errors.New("无效的令牌")
	ErrTokenExpired    = errors.New("令牌已过期")
	ErrPasswordFormat  = errors.New("密码格式错误")
)

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,len=32"`
	Email    string `json:"email" binding:"required,email"`
	Nickname string `json:"nickname" binding:"omitempty,max=100"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	TokenType    string    `json:"token_type"`
}

const BcryptCost = 12

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

	if len(input.Password) != 32 {
		return nil, ErrPasswordFormat
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), BcryptCost)
	if err != nil {
		return nil, err
	}

	nickname := input.Nickname
	if nickname == "" {
		nickname = input.Username
	}

	userData := &models.User{
		Username: input.Username,
		Password: string(hashedPassword),
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

func isEmail(account string) bool {
	return strings.Contains(account, "@")
}

func (s *UserService) LoginByAccount(account string, md5Password string, ip string) (*models.User, *TokenResponse, error) {
	var userData *models.User
	var err error

	if isEmail(account) {
		userData, err = s.repo.FindByEmail(account)
	} else {
		userData, err = s.repo.FindByUsername(account)
	}

	if err != nil {
		return nil, nil, ErrUserNotFound
	}

	if len(md5Password) != 32 {
		return nil, nil, ErrPasswordFormat
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(md5Password))
	if err != nil {
		return nil, nil, ErrInvalidPassword
	}

	if userData.Status == models.UserStatusDisabled {
		return nil, nil, ErrUserDisabled
	}

	tokenResponse, err := s.generateTokens(userData.ID, ip)
	if err != nil {
		return nil, nil, err
	}

	return userData, tokenResponse, nil
}

// GetProfile 获取用户资料
func (s *UserService) GetProfile(userID uint64) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) Login(username string, md5Password string, ip string) (*models.User, *TokenResponse, error) {
	userData, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, nil, ErrUserNotFound
	}

	if len(md5Password) != 32 {
		return nil, nil, ErrPasswordFormat
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(md5Password))
	if err != nil {
		return nil, nil, ErrInvalidPassword
	}

	if userData.Status == models.UserStatusDisabled {
		return nil, nil, ErrUserDisabled
	}

	tokenResponse, err := s.generateTokens(userData.ID, ip)
	if err != nil {
		return nil, nil, err
	}

	return userData, tokenResponse, nil
}

func (s *UserService) generateTokens(userID uint64, ip string) (*TokenResponse, error) {
	accessToken, accessExpiresAt, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	// 删除该用户的所有旧 refresh token，确保每个用户只有一个有效的 refresh token
	if err := s.repo.RevokeAllUserTokens(userID, models.TokenTypeRefresh); err != nil {
		return nil, err
	}

	refreshTokenRecord := &models.UserToken{
		UserID:    userID,
		Token:     utils.HashToken(refreshToken),
		Type:      models.TokenTypeRefresh,
		ExpiresAt: refreshExpiresAt,
	}

	if err := s.repo.CreateRefreshTokenAndUpdateLoginInfo(userID, ip, refreshTokenRecord); err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExpiresAt,
		TokenType:    "Bearer",
	}, nil
}

func (s *UserService) RefreshToken(refreshToken string) (*TokenResponse, error) {
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		if err == utils.ErrExpiredToken {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	tokenHash := utils.HashToken(refreshToken)
	token, err := s.repo.FindTokenByHash(tokenHash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if token.Type != models.TokenTypeRefresh {
		return nil, ErrInvalidToken
	}

	if err := s.repo.DeleteTokenByID(token.ID); err != nil {
		return nil, err
	}

	accessToken, accessExpiresAt, err := utils.GenerateAccessToken(claims.UserID)
	if err != nil {
		return nil, err
	}

	newRefreshToken, refreshExpiresAt, err := utils.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return nil, err
	}

	refreshTokenRecord := &models.UserToken{
		UserID:    claims.UserID,
		Token:     utils.HashToken(newRefreshToken),
		Type:      models.TokenTypeRefresh,
		ExpiresAt: refreshExpiresAt,
	}

	if err := s.repo.CreateToken(refreshTokenRecord); err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    accessExpiresAt,
		TokenType:    "Bearer",
	}, nil
}

func (s *UserService) Logout(accessToken string) error {
	return nil
}

func (s *UserService) LogoutByRefreshToken(refreshToken string) error {
	tokenHash := utils.HashToken(refreshToken)
	return s.repo.DeleteTokenByHash(tokenHash)
}

func (s *UserService) RevokeAllUserSessions(userID uint64) error {
	return s.repo.RevokeAllUserTokens(userID, models.TokenTypeRefresh)
}
