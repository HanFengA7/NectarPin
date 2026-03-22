// Package article 提供文章数据访问层
// 列表查询仅选 list 所需字段并分页，避免大字段与 N+1，保证高性能
package article

import (
	"database/sql"
	"strings"

	"nectarpin/api/models"

	"gorm.io/gorm"
)

// ListFields 列表项所需字段（不含 content，减轻传输与内存）
var ListFields = []string{"id", "author_id", "category_id", "title", "slug", "summary", "cover_image", "status", "view_count", "published_at", "created_at", "updated_at"}

func listSelectWithEffectiveViewCount() string {
	parts := make([]string, 0, len(ListFields))
	for _, f := range ListFields {
		if f == "view_count" {
			parts = append(parts, "COALESCE(article_view_stats.view_count, articles.view_count) AS view_count")
		} else {
			parts = append(parts, "articles."+f)
		}
	}
	return strings.Join(parts, ", ")
}

// ArticleRepository 文章仓储
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建文章仓储
func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// Create 创建文章
func (r *ArticleRepository) Create(m *models.Article) error {
	return r.db.Create(m).Error
}

// FindByID 按主键查询（完整记录，含 content）
func (r *ArticleRepository) FindByID(id uint64) (*models.Article, error) {
	var a models.Article
	err := r.db.First(&a, id).Error
	if err != nil {
		return nil, err
	}
	if err := r.hydrateEffectiveViewCount(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

// FindBySlug 按 slug 查询（完整记录）
func (r *ArticleRepository) FindBySlug(slug string) (*models.Article, error) {
	var a models.Article
	err := r.db.Where("slug = ?", slug).First(&a).Error
	if err != nil {
		return nil, err
	}
	if err := r.hydrateEffectiveViewCount(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

// ListFilter 列表过滤条件
type ListFilter struct {
	AuthorID *uint64 // 按作者筛选
	Status   *int16  // 按状态筛选，不传则不过滤
	Page     int     // 页码，从 1 开始
	PageSize int     // 每页条数，建议上限 100
}

// ListResult 分页列表结果
type ListResult struct {
	Items []models.Article
	Total int64
}

// List 分页列表，仅查询 ListFields，按 created_at 倒序
func (r *ArticleRepository) List(f ListFilter) (ListResult, error) {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.PageSize < 1 {
		f.PageSize = 20
	}
	if f.PageSize > 100 {
		f.PageSize = 100
	}

	query := r.db.Model(&models.Article{}).
		Select(listSelectWithEffectiveViewCount()).
		Joins("LEFT JOIN article_view_stats ON article_view_stats.article_id = articles.id")
	if f.AuthorID != nil {
		query = query.Where("author_id = ?", *f.AuthorID)
	}
	if f.Status != nil {
		query = query.Where("status = ?", *f.Status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return ListResult{}, err
	}

	var items []models.Article
	offset := (f.Page - 1) * f.PageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(f.PageSize).Find(&items).Error
	if err != nil {
		return ListResult{}, err
	}
	return ListResult{Items: items, Total: total}, nil
}

// IncrementViewCount 阅读量 +1：写入窄表 article_view_stats，避免 UPDATE articles 触发行重写（含大 content 时易 SLOW SQL）
func (r *ArticleRepository) IncrementViewCount(id uint64) error {
	return r.db.Exec(`
INSERT INTO article_view_stats (article_id, view_count)
SELECT a.id, a.view_count + 1
FROM articles a
WHERE a.id = ? AND a.deleted_at IS NULL
ON CONFLICT (article_id) DO UPDATE
SET view_count = article_view_stats.view_count + 1
`, id).Error
}

func (r *ArticleRepository) hydrateEffectiveViewCount(a *models.Article) error {
	var count sql.NullInt64
	err := r.db.Raw(
		`SELECT view_count FROM article_view_stats WHERE article_id = ? LIMIT 1`,
		a.ID,
	).Scan(&count).Error
	if err != nil {
		return err
	}
	if count.Valid {
		a.ViewCount = int(count.Int64)
	}
	return nil
}

// Update 按 ID 更新文章（只更新非零值字段）
func (r *ArticleRepository) Update(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&models.Article{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 软删除文章
func (r *ArticleRepository) Delete(id uint64) error {
	return r.db.Delete(&models.Article{}, id).Error
}
