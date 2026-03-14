// Package user 提供用户相关的数据模型定义
// 包含用户实体、用户令牌等核心数据结构
package user

import (
	"time"
)

// UserToken 用户令牌模型
// 用于存储用户的访问令牌、刷新令牌等认证信息
type UserToken struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`                   // 令牌唯一标识
	UserID    uint64    `gorm:"not null;index" json:"user_id"`          // 关联的用户 ID
	Token     string    `gorm:"size:500;not null;uniqueIndex" json:"-"` // 令牌字符串
	Type      int16     `gorm:"not null" json:"type"`                   // 令牌类型
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`             // 过期时间
	Revoked   bool      `gorm:"not null;default:false" json:"revoked"`  // 是否已撤销
	CreatedAt time.Time `gorm:"not null" json:"created_at"`             // 创建时间
}

// TableName 指定用户令牌表名称
// 返回:
//   - string: 数据库表名 "user_tokens"
func (UserToken) TableName() string {
	return "user_tokens"
}

// 令牌类型常量
const (
	TokenTypeAccess  int16 = 1 // 访问令牌
	TokenTypeRefresh int16 = 2 // 刷新令牌
	TokenTypeAPI     int16 = 3 // API 令牌
)
