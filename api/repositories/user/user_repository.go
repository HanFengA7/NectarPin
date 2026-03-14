// Package user 提供用户数据访问层实现
// 封装用户相关的数据库操作，包括 CRUD 和查询功能
package user

import (
	"time"

	"nectarpin/api/models"

	"gorm.io/gorm"
)

// UserRepository 用户数据仓库
// 提供用户数据的持久化操作接口
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户数据仓库实例
// 参数:
//   - db: GORM 数据库连接实例
//
// 返回:
//   - *UserRepository: 用户数据仓库实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建新用户
// 参数:
//   - userData: 用户数据实体
//
// 返回:
//   - error: 创建过程中的错误信息
func (r *UserRepository) Create(userData *models.User) error {
	return r.db.Create(userData).Error
}

// FindByID 根据 ID 查询用户
// 参数:
//   - id: 用户唯一标识
//
// 返回:
//   - *models.User: 查询到的用户数据
//   - error: 查询过程中的错误信息
func (r *UserRepository) FindByID(id uint64) (*models.User, error) {
	var userData models.User
	err := r.db.First(&userData, id).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

// FindByUsername 根据用户名查询用户
// 参数:
//   - username: 用户名
//
// 返回:
//   - *models.User: 查询到的用户数据
//   - error: 查询过程中的错误信息
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var userData models.User
	err := r.db.Where("username = ?", username).First(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

// FindByEmail 根据邮箱查询用户
// 参数:
//   - email: 邮箱地址
//
// 返回:
//   - *models.User: 查询到的用户数据
//   - error: 查询过程中的错误信息
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var userData models.User
	err := r.db.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

// ExistsByUsername 检查用户名是否已存在
// 参数:
//   - username: 用户名
//
// 返回:
//   - bool: true 表示已存在
//   - error: 查询过程中的错误信息
func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// ExistsByEmail 检查邮箱是否已存在
// 参数:
//   - email: 邮箱地址
//
// 返回:
//   - bool: true 表示已存在
//   - error: 查询过程中的错误信息
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// UpdateLoginInfo 更新用户登录信息
// 参数:
//   - id: 用户唯一标识
//   - ip: 登录 IP 地址
//
// 返回:
//   - error: 更新过程中的错误信息
//
// 该方法会更新最后登录时间、登录 IP 和登录次数
func (r *UserRepository) UpdateLoginInfo(id uint64, ip string) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_login_at": now,
		"last_login_ip": ip,
		"login_count":   gorm.Expr("login_count + 1"),
	}).Error
}

// CreateToken 创建用户令牌
// 参数:
//   - token: 令牌数据实体
//
// 返回:
//   - error: 创建过程中的错误信息
func (r *UserRepository) CreateToken(token *models.UserToken) error {
	return r.db.Create(token).Error
}

// FindTokenByToken 根据令牌字符串查询令牌
// 参数:
//   - tokenString: 令牌字符串
//
// 返回:
//   - *models.UserToken: 查询到的令牌数据
//   - error: 查询过程中的错误信息
//
// 只查询未被撤销的令牌
func (r *UserRepository) FindTokenByToken(tokenString string) (*models.UserToken, error) {
	var token models.UserToken
	err := r.db.Where("token = ? AND revoked = ?", tokenString, false).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// RevokeToken 撤销指定令牌
// 参数:
//   - tokenString: 令牌字符串
//
// 返回:
//   - error: 撤销过程中的错误信息
func (r *UserRepository) RevokeToken(tokenString string) error {
	return r.db.Model(&models.UserToken{}).Where("token = ?", tokenString).Update("revoked", true).Error
}

// RevokeAllUserTokens 撤销用户指定类型的所有令牌
// 参数:
//   - userID: 用户唯一标识
//   - tokenType: 令牌类型
//
// 返回:
//   - error: 撤销过程中的错误信息
func (r *UserRepository) RevokeAllUserTokens(userID uint64, tokenType int16) error {
	return r.db.Model(&models.UserToken{}).Where("user_id = ? AND type = ?", userID, tokenType).Update("revoked", true).Error
}
