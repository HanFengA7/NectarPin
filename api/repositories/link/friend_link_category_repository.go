package link

import (
	linkmodel "nectarpin/api/models/link"

	"gorm.io/gorm"
)

type FriendLinkCategoryRepository struct {
	db *gorm.DB
}

func NewFriendLinkCategoryRepository(db *gorm.DB) *FriendLinkCategoryRepository {
	return &FriendLinkCategoryRepository{db: db}
}

func (r *FriendLinkCategoryRepository) Create(m *linkmodel.FriendLinkCategory) error {
	return r.db.Create(m).Error
}

func (r *FriendLinkCategoryRepository) FindByID(id uint64) (*linkmodel.FriendLinkCategory, error) {
	var c linkmodel.FriendLinkCategory
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FriendLinkCategoryRepository) FindByName(name string) (*linkmodel.FriendLinkCategory, error) {
	var c linkmodel.FriendLinkCategory
	err := r.db.Where("name = ?", name).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FriendLinkCategoryRepository) FindBySlug(slug string) (*linkmodel.FriendLinkCategory, error) {
	var c linkmodel.FriendLinkCategory
	err := r.db.Where("slug = ?", slug).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FriendLinkCategoryRepository) ListAll() ([]linkmodel.FriendLinkCategory, error) {
	var items []linkmodel.FriendLinkCategory
	err := r.db.Order("sort_order ASC, id ASC").Find(&items).Error
	return items, err
}

func (r *FriendLinkCategoryRepository) Update(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&linkmodel.FriendLinkCategory{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FriendLinkCategoryRepository) Delete(id uint64) error {
	return r.db.Delete(&linkmodel.FriendLinkCategory{}, id).Error
}

func (r *FriendLinkCategoryRepository) ClearCategoryOnLinks(categoryID uint64) error {
	return r.db.Model(&linkmodel.FriendLink{}).
		Where("category_id = ?", categoryID).
		Update("category_id", nil).Error
}
