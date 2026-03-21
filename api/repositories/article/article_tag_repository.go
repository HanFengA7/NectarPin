package article

import (
	"nectarpin/api/models"

	"gorm.io/gorm"
)

type ArticleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository(db *gorm.DB) *ArticleTagRepository {
	return &ArticleTagRepository{db: db}
}

func (r *ArticleTagRepository) Create(m *models.ArticleTag) error {
	return r.db.Create(m).Error
}

func (r *ArticleTagRepository) FindByID(id uint64) (*models.ArticleTag, error) {
	var t models.ArticleTag
	err := r.db.First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *ArticleTagRepository) FindByName(name string) (*models.ArticleTag, error) {
	var t models.ArticleTag
	err := r.db.Where("name = ?", name).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *ArticleTagRepository) FindBySlug(slug string) (*models.ArticleTag, error) {
	var t models.ArticleTag
	err := r.db.Where("slug = ?", slug).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// ListAll 全量返回，按 name ASC 排序
func (r *ArticleTagRepository) ListAll() ([]models.ArticleTag, error) {
	var items []models.ArticleTag
	err := r.db.Order("name ASC").Find(&items).Error
	return items, err
}

func (r *ArticleTagRepository) Update(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&models.ArticleTag{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ArticleTagRepository) Delete(id uint64) error {
	return r.db.Delete(&models.ArticleTag{}, id).Error
}

// IncrementArticleCount 原子更新 article_count
func (r *ArticleTagRepository) IncrementArticleCount(id uint64, delta int) error {
	return r.db.Model(&models.ArticleTag{}).
		Where("id = ?", id).
		UpdateColumn("article_count", gorm.Expr("article_count + ?", delta)).Error
}

// --- ArticleTagMapping 关联操作 ---

// GetTagIDsByArticleID 获取文章关联的所有标签 ID
func (r *ArticleTagRepository) GetTagIDsByArticleID(articleID uint64) ([]uint64, error) {
	var ids []uint64
	err := r.db.Model(&models.ArticleTagMapping{}).
		Where("article_id = ?", articleID).
		Pluck("tag_id", &ids).Error
	return ids, err
}

// ReplaceArticleTags 替换文章的标签关联（一次 DELETE + 批量 INSERT）
// 返回新增和移除的标签 ID，用于维护计数器
func (r *ArticleTagRepository) ReplaceArticleTags(tx *gorm.DB, articleID uint64, newTagIDs []uint64) (added, removed []uint64, err error) {
	var oldIDs []uint64
	if err = tx.Model(&models.ArticleTagMapping{}).
		Where("article_id = ?", articleID).
		Pluck("tag_id", &oldIDs).Error; err != nil {
		return
	}

	oldSet := make(map[uint64]bool, len(oldIDs))
	for _, id := range oldIDs {
		oldSet[id] = true
	}
	newSet := make(map[uint64]bool, len(newTagIDs))
	for _, id := range newTagIDs {
		newSet[id] = true
	}

	for _, id := range newTagIDs {
		if !oldSet[id] {
			added = append(added, id)
		}
	}
	for _, id := range oldIDs {
		if !newSet[id] {
			removed = append(removed, id)
		}
	}

	if err = tx.Where("article_id = ?", articleID).Delete(&models.ArticleTagMapping{}).Error; err != nil {
		return
	}

	if len(newTagIDs) > 0 {
		mappings := make([]models.ArticleTagMapping, len(newTagIDs))
		for i, tagID := range newTagIDs {
			mappings[i] = models.ArticleTagMapping{ArticleID: articleID, TagID: tagID}
		}
		err = tx.Create(&mappings).Error
	}
	return
}

// DeleteArticleTagMappings 删除文章的所有标签关联
func (r *ArticleTagRepository) DeleteArticleTagMappings(articleID uint64) error {
	return r.db.Where("article_id = ?", articleID).Delete(&models.ArticleTagMapping{}).Error
}

// DeleteTagMappings 删除某标签的所有关联记录
func (r *ArticleTagRepository) DeleteTagMappings(tagID uint64) error {
	return r.db.Where("tag_id = ?", tagID).Delete(&models.ArticleTagMapping{}).Error
}

// DB 返回底层 gorm.DB，用于事务
func (r *ArticleTagRepository) DB() *gorm.DB {
	return r.db
}
