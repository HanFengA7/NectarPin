package link

import "time"

// FriendLink 单条友链
type FriendLink struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	CategoryID  *uint64   `gorm:"index:idx_friend_links_category_sort" json:"category_id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	URL         string    `gorm:"size:500;not null" json:"url"`
	Description string    `gorm:"type:text" json:"description"`
	AvatarURL   string    `gorm:"size:500" json:"avatar_url"`
	SortOrder   int       `gorm:"not null;default:0;index:idx_friend_links_category_sort" json:"sort_order"`
	IsEnabled   bool      `gorm:"not null;default:true" json:"is_enabled"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

func (FriendLink) TableName() string {
	return "friend_links"
}
