package models

import (
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

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
