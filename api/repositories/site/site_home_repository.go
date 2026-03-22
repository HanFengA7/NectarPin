package site

import (
	"errors"

	sitemodel "nectarpin/api/models/site"

	"gorm.io/gorm"
)

// SiteHomeRepository 首页配置仓储
type SiteHomeRepository struct {
	db *gorm.DB
}

// NewSiteHomeRepository 创建仓储
func NewSiteHomeRepository(db *gorm.DB) *SiteHomeRepository {
	return &SiteHomeRepository{db: db}
}

// GetSingleton 按主键读取单行配置（无记录时返回 nil, nil）
func (r *SiteHomeRepository) GetSingleton() (*sitemodel.SiteHomeConfig, error) {
	var row sitemodel.SiteHomeConfig
	err := r.db.Where("id = ?", sitemodel.SiteHomeSingletonID).First(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

// UpsertSingleton 写入或更新单行（Save 按主键 upsert）
func (r *SiteHomeRepository) UpsertSingleton(payload string) error {
	row := sitemodel.SiteHomeConfig{
		ID:      sitemodel.SiteHomeSingletonID,
		Payload: payload,
	}
	return r.db.Save(&row).Error
}
