package models

import (
	"nectarpin/internal/database"
	"nectarpin/internal/utils"
	"time"

	"gorm.io/gorm"
)

/*
User 用户表

	ID 用户ID
	UserName 用户名
	NickName 昵称
	Password 密码
	Email 邮箱
	Avatar 头像
	Role 角色(0:普通用户 1:管理员)
	CreatedAt 创建时间
	UpdatedAt 更新时间
	DeletedAt 删除时间
*/
type User struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	UserName  string         `json:"user_name"`
	NickName  string         `json:"nick_name"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Avatar    string         `json:"avatar"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

/*
CreateUser 创建用户 [0.1][2025.12.06]

	@param user *UserModel
	@return msg string 提示信息
	@return code int 状态码
*/
func CreateUser(user *User) (msg string, code int) {
	db := database.GetDB()
	var err error
	//[1]数据校验
	//[2]密码加密
	msg, code, user.Password = utils.EncryptUserPassword(user.Password)
	if code != 200 {
		return msg, 500
	}
	//[3]写入数据库
	err = db.Create(&user).Error
	if err != nil {
		return "创建用户失败", 500
	}
	return "创建用户成功", 200
}

/*
ExistUserOrEmail 检测(用户/邮箱)是否存在 [0.2][2025.12.07]

	@param userName string 用户名(可选)
	@param email string 邮箱(可选)
	@return msg string 提示信息
	@return code int 状态码
	@return bool 检测结果
*/
func ExistUserOrEmail(userName string, email string) (msg string, code int, exist bool) {
	db := database.GetDB()
	var user User

	// 检查用户名是否存在
	if userName != "" {
		usernameRows := db.Where("user_name = ?", userName).First(&user).RowsAffected
		if usernameRows > 0 {
			// 如果同时传了邮箱，且邮箱属于该用户，则跳过邮箱检测
			if email != "" && user.Email == email {
				return "检测完成,无重复!", 200, false
			}
			return "用户名已存在!", 2001, true
		}
	}

	// 检查邮箱是否存在
	if email != "" {
		var emailUser User
		emailRows := db.Where("email = ?", email).First(&emailUser).RowsAffected
		if emailRows > 0 {
			return "邮箱已存在!", 2002, true
		}
	}

	return "检测完成,无重复!", 200, false
}
