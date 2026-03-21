// Package article 提供文章相关的数据模型定义
package article

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章实体模型
// 设计要点：author_id/status/published_at/slug 建索引以支持高性能列表与单条查询
type Article struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	AuthorID    uint64         `gorm:"not null;index:idx_articles_author_id" json:"author_id"`
	CategoryID  *uint64        `gorm:"index:idx_articles_category_id" json:"category_id"`
	Title       string         `gorm:"size:200;not null" json:"title"`
	Slug        string         `gorm:"size:220;not null;uniqueIndex:idx_articles_slug" json:"slug"`
	Summary     string         `gorm:"type:text" json:"summary"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	CoverImage  string         `gorm:"size:500" json:"cover_image"`
	Status      int16          `gorm:"not null;default:0;index:idx_articles_status" json:"status"`
	ViewCount   int            `gorm:"not null;default:0" json:"view_count"`
	PublishedAt *time.Time     `gorm:"index:idx_articles_published_at" json:"published_at"`
	CreatedAt   time.Time      `gorm:"not null;index:idx_articles_created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定文章表名称
func (Article) TableName() string {
	return "articles"
}

// 文章状态常量
const (
	ArticleStatusDraft    int16 = 0 // 草稿
	ArticleStatusPublished int16 = 1 // 已发布
	ArticleStatusArchived int16 = 2 // 已归档
)
