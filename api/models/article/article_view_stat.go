package article

// ArticleViewStat 文章阅读量窄表：仅 article_id + view_count，避免每次 +1 都改写 articles 大正文行（PG 上易慢）
type ArticleViewStat struct {
	ArticleID uint64 `gorm:"primaryKey;column:article_id" json:"article_id"`
	ViewCount int    `gorm:"not null;default:0" json:"view_count"`
}

// TableName 指定表名
func (ArticleViewStat) TableName() string {
	return "article_view_stats"
}
