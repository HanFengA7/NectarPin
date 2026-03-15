// Package user 提供用户相关的数据模型定义
// 包含用户实体、用户令牌等核心数据结构
package user

import (
	"time"

	"gorm.io/gorm"
)

// User 用户实体模型
// 存储用户的基本信息、状态和认证相关数据
type User struct {
	ID               uint64         `gorm:"primaryKey" json:"id"`                             // 用户唯一标识
	Username         string         `gorm:"size:50;not null;uniqueIndex" json:"username"`     // 用户名，唯一
	Password         string         `gorm:"size:60;not null" json:"-"`                        // 密码（bcrypt加密），不返回给前端
	Email            string         `gorm:"size:255;not null;uniqueIndex" json:"email"`       // 邮箱地址，唯一
	Nickname         string         `gorm:"size:100" json:"nickname"`                         // 昵称
	Avatar           string         `gorm:"size:500" json:"avatar"`                           // 头像 URL
	Bio              string         `gorm:"type:text" json:"bio"`                             // 个人简介
	Status           int16          `gorm:"not null;default:1" json:"status"`                 // 账户状态
	Role             int16          `gorm:"not null;default:1" json:"role"`                   // 用户角色
	EmailVerified    bool           `gorm:"not null;default:false" json:"email_verified"`     // 邮箱是否已验证
	EmailVerifiedAt  *time.Time     `json:"email_verified_at"`                                // 邮箱验证时间
	TwoFactorEnabled bool           `gorm:"not null;default:false" json:"two_factor_enabled"` // 是否启用双因素认证
	TwoFactorSecret  string         `gorm:"size:255" json:"-"`                                // 双因素认证密钥
	LastLoginAt      *time.Time     `json:"last_login_at"`                                    // 最后登录时间
	LastLoginIP      string         `gorm:"size:45" json:"last_login_ip"`                     // 最后登录 IP
	LoginCount       int            `gorm:"not null;default:0" json:"login_count"`            // 登录次数
	CreatedAt        time.Time      `gorm:"not null" json:"created_at"`                       // 创建时间
	UpdatedAt        time.Time      `gorm:"not null" json:"updated_at"`                       // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`                                   // 软删除时间
}

// TableName 指定用户表名称
// 返回:
//   - string: 数据库表名 "users"
func (User) TableName() string {
	return "users"
}

// 用户状态常量
const (
	UserStatusDisabled int16 = 0 // 已禁用
	UserStatusActive   int16 = 1 // 正常
	UserStatusPending  int16 = 2 // 待激活
)

// 用户角色常量
const (
	UserRoleUser  int16 = 1 // 普通用户
	UserRoleAdmin int16 = 2 // 管理员
)
