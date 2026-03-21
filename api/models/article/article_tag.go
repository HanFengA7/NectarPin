package article

import "time"

// ArticleTag 文章标签实体模型
type ArticleTag struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null;uniqueIndex:idx_article_tags_name" json:"name"`
	Slug         string    `gorm:"size:120;not null;uniqueIndex:idx_article_tags_slug" json:"slug"`
	ArticleCount int       `gorm:"not null;default:0" json:"article_count"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
}

func (ArticleTag) TableName() string {
	return "article_tags"
}

// ArticleTagMapping 文章-标签多对多关联表
type ArticleTagMapping struct {
	ArticleID uint64 `gorm:"primaryKey;index:idx_tag_article,priority:1" json:"article_id"`
	TagID     uint64 `gorm:"primaryKey;index:idx_tag_article,priority:2" json:"tag_id"`
}

func (ArticleTagMapping) TableName() string {
	return "article_tag_mappings"
}
