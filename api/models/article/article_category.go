package article

import "time"

// ArticleCategory 文章分类实体模型
type ArticleCategory struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null;uniqueIndex:idx_article_categories_name" json:"name"`
	Slug         string    `gorm:"size:120;not null;uniqueIndex:idx_article_categories_slug" json:"slug"`
	Description  string    `gorm:"type:text" json:"description"`
	SortOrder    int       `gorm:"not null;default:0;index:idx_article_categories_sort_order" json:"sort_order"`
	ArticleCount int       `gorm:"not null;default:0" json:"article_count"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
}

func (ArticleCategory) TableName() string {
	return "article_categories"
}
