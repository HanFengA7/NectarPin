package link

import (
	linkmodel "nectarpin/api/models/link"

	"gorm.io/gorm"
)

type FriendLinkRepository struct {
	db *gorm.DB
}

func NewFriendLinkRepository(db *gorm.DB) *FriendLinkRepository {
	return &FriendLinkRepository{db: db}
}

func (r *FriendLinkRepository) Create(m *linkmodel.FriendLink) error {
	return r.db.Create(m).Error
}

func (r *FriendLinkRepository) FindByID(id uint64) (*linkmodel.FriendLink, error) {
	var m linkmodel.FriendLink
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *FriendLinkRepository) ListAllForAdmin() ([]linkmodel.FriendLink, error) {
	var items []linkmodel.FriendLink
	err := r.db.Order("sort_order ASC, id ASC").Find(&items).Error
	return items, err
}

func (r *FriendLinkRepository) ListEnabledPublic() ([]linkmodel.FriendLink, error) {
	var items []linkmodel.FriendLink
	err := r.db.Where("is_enabled = ?", true).
		Order("sort_order ASC, id ASC").
		Find(&items).Error
	return items, err
}

func (r *FriendLinkRepository) Update(id uint64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.Model(&linkmodel.FriendLink{}).Where("id = ?", id).Updates(updates).Error
}

func (r *FriendLinkRepository) Delete(id uint64) error {
	return r.db.Delete(&linkmodel.FriendLink{}, id).Error
}
