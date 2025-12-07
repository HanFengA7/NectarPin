package repository

import (
	"nectarpin/internal/models"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问层
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建新的用户 Repository 实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

/*
Create 创建用户记录

	@param user *models.User
	@return error
*/
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

/*
ExistsByUsername 检查用户名是否存在

	@param userName string
	@return *models.User 如果存在返回用户信息，否则返回 nil
	@return error
*/
func (r *UserRepository) ExistsByUsername(userName string) (*models.User, error) {
	var user models.User
	result := r.db.Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

/*
ExistsByEmail 检查邮箱是否存在

	@param email string
	@return bool 是否存在
	@return error
*/
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

/*
FindByID 根据ID查找用户

	@param id uint
	@return *models.User
	@return error
*/
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

/*
FindByUsername 根据用户名查找用户

	@param userName string
	@return *models.User
	@return error
*/
func (r *UserRepository) FindByUsername(userName string) (*models.User, error) {
	var user models.User
	result := r.db.Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

/*
Update 更新用户信息

	@param user *models.User
	@return error
*/
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

/*
Delete 删除用户

	@param id uint
	@return error
*/
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
