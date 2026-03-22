package link

import (
	"errors"

	linkmodel "nectarpin/api/models/link"

	"gorm.io/gorm"
)

const friendLinkPageSingletonID uint64 = 1

type FriendLinkPageRepository struct {
	db *gorm.DB
}

func NewFriendLinkPageRepository(db *gorm.DB) *FriendLinkPageRepository {
	return &FriendLinkPageRepository{db: db}
}

// GetOrCreateSingleton 返回 id=1 的友链页配置，不存在则插入空行
func (r *FriendLinkPageRepository) GetOrCreateSingleton() (*linkmodel.FriendLinkPage, error) {
	var p linkmodel.FriendLinkPage
	err := r.db.First(&p, friendLinkPageSingletonID).Error
	if err == nil {
		return &p, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	p = linkmodel.FriendLinkPage{
		ID:        friendLinkPageSingletonID,
		IntroHTML: "",
	}
	if err := r.db.Create(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *FriendLinkPageRepository) UpdateIntroHTML(html string) error {
	_, err := r.GetOrCreateSingleton()
	if err != nil {
		return err
	}
	return r.db.Model(&linkmodel.FriendLinkPage{}).
		Where("id = ?", friendLinkPageSingletonID).
		Update("intro_html", html).Error
}
