package link

import "time"

// FriendLinkPage 友链页单页配置（固定主键 1：页顶简介等）
type FriendLinkPage struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	IntroHTML string    `gorm:"type:text" json:"intro_html"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (FriendLinkPage) TableName() string {
	return "friend_link_page"
}
