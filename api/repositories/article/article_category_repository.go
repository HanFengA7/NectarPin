package article

import (
	"nectarpin/api/models"

	"gorm.io/gorm"
)

type ArticleCategoryRepository struct {
	db *gorm.DB
}

func NewArticleCategoryRepository(db *gorm.DB) *ArticleCategoryRepository {
	return &ArticleCategoryRepository{db: db}
}

func (r *ArticleCategoryRepository) Create(m *models.ArticleCategory) error {
	return r.db.Create(m).Error
}

func (r *ArticleCategoryRepository) FindByID(id uint64) (*models.ArticleCategory, error) {
	var c models.ArticleCategory
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ArticleCategoryRepository) FindByName(name string) (*models.ArticleCategory, error) {
	var c models.ArticleCategory
	err := r.db.Where("name = ?", name).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ArticleCategoryRepository) FindBySlug(slug string) (*models.ArticleCategory, error) {
	var c models.ArticleCategory
	err := r.db.Where("slug = ?", slug).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// ListAll 全量返回，按 sort_order ASC, id ASC 排序
func (r *ArticleCategoryRepository) ListAll() ([]models.ArticleCategory, error) {
	var items []models.ArticleCategory
	err := r.db.Order("sort_order ASC, id ASC").Find(&items).Error
	return items, err
}

func (r *ArticleCategoryRepository) Update(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&models.ArticleCategory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ArticleCategoryRepository) Delete(id uint64) error {
	return r.db.Delete(&models.ArticleCategory{}, id).Error
}

// IncrementArticleCount 原子自增 article_count
func (r *ArticleCategoryRepository) IncrementArticleCount(id uint64, delta int) error {
	return r.db.Model(&models.ArticleCategory{}).
		Where("id = ?", id).
		UpdateColumn("article_count", gorm.Expr("article_count + ?", delta)).Error
}

// ClearArticleCategory 将某分类下的所有文章 category_id 置 NULL（删除分类时使用）
func (r *ArticleCategoryRepository) ClearArticleCategory(categoryID uint64) error {
	return r.db.Table("articles").
		Where("category_id = ?", categoryID).
		Update("category_id", nil).Error
}
