package user

import (
	"errors"
	"strings"
	"time"

	"nectarpin/api/models"
	userrepo "nectarpin/api/repositories/user"
	"nectarpin/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *userrepo.UserRepository
}

func NewUserService(repo *userrepo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var (
	ErrUsernameExists    = errors.New("用户名已存在")
	ErrEmailExists       = errors.New("邮箱已存在")
	ErrUserNotFound      = errors.New("用户不存在")
	ErrUserDisabled      = errors.New("用户已被禁用")
	ErrInvalidPassword   = errors.New("密码错误")
	ErrInvalidToken      = errors.New("无效的令牌")
	ErrTokenExpired      = errors.New("令牌已过期")
	ErrPasswordFormat    = errors.New("密码格式错误")
	ErrPasswordUnchanged = errors.New("新密码不能与当前密码相同")
	ErrSessionNotFound   = errors.New("会话不存在")
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

// UpdateProfileInput 更新当前用户资料（邮箱必填，与 OpenAPI 一致）
type UpdateProfileInput struct {
	Nickname string `json:"nickname" binding:"max=100"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar" binding:"max=500"`
}

// UpdateProfile 更新用户昵称、邮箱、头像
func (s *UserService) UpdateProfile(userID uint64, input *UpdateProfileInput) (*models.User, error) {
	_, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	email := strings.TrimSpace(input.Email)
	nickname := strings.TrimSpace(input.Nickname)
	avatar := strings.TrimSpace(input.Avatar)

	exists, err := s.repo.ExistsByEmailExceptUser(email, userID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	updates := map[string]interface{}{
		"email":    email,
		"nickname": nickname,
		"avatar":   avatar,
	}
	if err := s.repo.UpdateUserFields(userID, updates); err != nil {
		return nil, err
	}

	return s.repo.FindByID(userID)
}

// ChangePassword 校验旧密码后更新；oldMD5、newMD5 均为前端 MD5(明文) 的 32 位十六进制，与登录接口一致
func (s *UserService) ChangePassword(userID uint64, oldMD5, newMD5 string) error {
	if len(oldMD5) != 32 || len(newMD5) != 32 {
		return ErrPasswordFormat
	}
	if oldMD5 == newMD5 {
		return ErrPasswordUnchanged
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldMD5))
	if err != nil {
		return ErrInvalidPassword
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newMD5), BcryptCost)
	if err != nil {
		return err
	}

	return s.repo.UpdateUserFields(userID, map[string]interface{}{
		"password": string(hashed),
	})
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

	// 允许多端同时登录：保留已有 refresh 会话，仅新增本条
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

// UserSessionItem 返回给前端的登录会话（不含令牌明文或 hash）
type UserSessionItem struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	IsCurrent bool      `json:"is_current"`
}

// ListUserSessions 列出 refresh 会话；若传入当前 refresh 明文则可标记 is_current
func (s *UserService) ListUserSessions(userID uint64, optionalRefreshPlain string) ([]UserSessionItem, error) {
	tokens, err := s.repo.ListRefreshTokensByUserID(userID)
	if err != nil {
		return nil, err
	}

	var keepHash string
	if optionalRefreshPlain != "" {
		keepHash = utils.HashToken(optionalRefreshPlain)
	}

	items := make([]UserSessionItem, 0, len(tokens))
	for _, t := range tokens {
		items = append(items, UserSessionItem{
			ID:        t.ID,
			CreatedAt: t.CreatedAt,
			ExpiresAt: t.ExpiresAt,
			IsCurrent: keepHash != "" && t.Token == keepHash,
		})
	}
	return items, nil
}

// RevokeSessionByID 撤销指定 refresh 会话（踢下线该客户端）
func (s *UserService) RevokeSessionByID(userID, sessionID uint64) error {
	err := s.repo.DeleteRefreshTokenByIDAndUserID(sessionID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrSessionNotFound
		}
		return err
	}
	return nil
}

// RevokeOtherSessions 撤销除当前 refresh 外的全部会话（需传入本机 refresh 以校验）
func (s *UserService) RevokeOtherSessions(userID uint64, currentRefreshPlain string) error {
	if currentRefreshPlain == "" {
		return ErrInvalidToken
	}
	hash := utils.HashToken(currentRefreshPlain)
	token, err := s.repo.FindTokenByHash(hash)
	if err != nil || token.UserID != userID || token.Type != models.TokenTypeRefresh {
		return ErrInvalidToken
	}
	return s.repo.RevokeRefreshTokensExceptHash(userID, hash)
}
