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
