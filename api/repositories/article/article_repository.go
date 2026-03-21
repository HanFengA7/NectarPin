// Package article 提供文章数据访问层
// 列表查询仅选 list 所需字段并分页，避免大字段与 N+1，保证高性能
package article

import (
	"nectarpin/api/models"

	"gorm.io/gorm"
)

// ListFields 列表项所需字段（不含 content，减轻传输与内存）
var ListFields = []string{"id", "author_id", "category_id", "title", "slug", "summary", "cover_image", "status", "view_count", "published_at", "created_at", "updated_at"}

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
	return &a, nil
}

// FindBySlug 按 slug 查询（完整记录）
func (r *ArticleRepository) FindBySlug(slug string) (*models.Article, error) {
	var a models.Article
	err := r.db.Where("slug = ?", slug).First(&a).Error
	if err != nil {
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

	query := r.db.Model(&models.Article{}).Select(ListFields)
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

// IncrementViewCount 阅读量 +1（原子自增，避免并发竞态）
func (r *ArticleRepository) IncrementViewCount(id uint64) error {
	return r.db.Model(&models.Article{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
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
