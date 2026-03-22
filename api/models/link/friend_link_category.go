package link

import "time"

// FriendLinkCategory 友链分组（后台「页面设置」管理）
type FriendLinkCategory struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null;uniqueIndex:idx_friend_link_categories_name" json:"name"`
	Slug        string    `gorm:"size:120;not null;uniqueIndex:idx_friend_link_categories_slug" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	SortOrder   int       `gorm:"not null;default:0;index:idx_friend_link_categories_sort" json:"sort_order"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

func (FriendLinkCategory) TableName() string {
	return "friend_link_categories"
}
