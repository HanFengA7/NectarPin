package user

import (
	"time"
)

type UserToken struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null;index:idx_user_type" json:"user_id"`
	Token     string    `gorm:"size:500;not null;uniqueIndex" json:"-"`
	Type      int16     `gorm:"not null;index:idx_user_type" json:"type"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

const (
	TokenTypeAccess  int16 = 1
	TokenTypeRefresh int16 = 2
	TokenTypeAPI     int16 = 3
)
