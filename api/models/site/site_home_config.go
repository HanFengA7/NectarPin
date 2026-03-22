// Package site 站点相关数据模型
package site

import "time"

// SiteHomeSingletonID 首页配置单行记录主键（固定为 1，便于主键点查）
const SiteHomeSingletonID uint64 = 1

// SiteHomeConfig 首页站点配置（整页 JSON 存一列，避免多行键值扫描）
type SiteHomeConfig struct {
	ID        uint64    `gorm:"primaryKey"`
	Payload   string    `gorm:"type:text;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName 指定表名
func (SiteHomeConfig) TableName() string {
	return "site_home_config"
}
