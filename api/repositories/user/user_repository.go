package user

import (
	"time"

	"nectarpin/api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(userData *models.User) error {
	return r.db.Create(userData).Error
}

func (r *UserRepository) FindByID(id uint64) (*models.User, error) {
	var userData models.User
	err := r.db.First(&userData, id).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

// FindByIDWithFields 根据 ID 查找用户，返回指定字段
func (r *UserRepository) FindByIDWithFields(id uint64, fields ...string) (*models.User, error) {
	var userData models.User
	query := r.db.Select(fields)
	err := query.First(&userData, id).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var userData models.User
	err := r.db.Where("username = ?", username).First(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var userData models.User
	err := r.db.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// ExistsByEmailExceptUser 是否存在其他用户使用该邮箱
func (r *UserRepository) ExistsByEmailExceptUser(email string, exceptUserID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ? AND id <> ?", email, exceptUserID).Count(&count).Error
	return count > 0, err
}

// UpdateUserFields 按 map 更新用户字段（零值字段亦会写入）
func (r *UserRepository) UpdateUserFields(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) UpdateLoginInfo(id uint64, ip string) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_login_at": now,
		"last_login_ip": ip,
		"login_count":   gorm.Expr("login_count + 1"),
	}).Error
}

func (r *UserRepository) CreateToken(token *models.UserToken) error {
	return r.db.Create(token).Error
}

func (r *UserRepository) CreateRefreshTokenAndUpdateLoginInfo(userID uint64, ip string, refreshToken *models.UserToken) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now()

		if err := tx.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
			"last_login_at": now,
			"last_login_ip": ip,
			"login_count":   gorm.Expr("login_count + 1"),
			"updated_at":    now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(refreshToken).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepository) FindTokenByHash(tokenHash string) (*models.UserToken, error) {
	var token models.UserToken
	err := r.db.Where("token = ?", tokenHash).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *UserRepository) DeleteTokenByID(id uint64) error {
	return r.db.Delete(&models.UserToken{}, id).Error
}

func (r *UserRepository) DeleteTokenByHash(tokenHash string) error {
	return r.db.Where("token = ?", tokenHash).Delete(&models.UserToken{}).Error
}

func (r *UserRepository) RevokeAllUserTokens(userID uint64, tokenType int16) error {
	return r.db.Where("user_id = ? AND type = ?", userID, tokenType).Delete(&models.UserToken{}).Error
}

// ListRefreshTokensByUserID 列出用户全部 refresh 会话（按创建时间倒序）
func (r *UserRepository) ListRefreshTokensByUserID(userID uint64) ([]models.UserToken, error) {
	var tokens []models.UserToken
	err := r.db.Where("user_id = ? AND type = ?", userID, models.TokenTypeRefresh).
		Order("created_at DESC").
		Find(&tokens).Error
	return tokens, err
}

// DeleteRefreshTokenByIDAndUserID 删除指定 refresh 会话（校验归属）
func (r *UserRepository) DeleteRefreshTokenByIDAndUserID(id, userID uint64) error {
	res := r.db.Where("id = ? AND user_id = ? AND type = ?", id, userID, models.TokenTypeRefresh).
		Delete(&models.UserToken{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// RevokeRefreshTokensExceptHash 撤销除指定 hash 外的全部 refresh 会话（多端踢下线）
func (r *UserRepository) RevokeRefreshTokensExceptHash(userID uint64, keepHash string) error {
	return r.db.Where("user_id = ? AND type = ? AND token <> ?", userID, models.TokenTypeRefresh, keepHash).
		Delete(&models.UserToken{}).Error
}
