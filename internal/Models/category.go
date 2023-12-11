package Models

import (
	"gorm.io/gorm"
	"time"
)

/*
Category [分类表结构][231209][0.1]

	ID				uint			分类ID
	CreatedAt		time.Time		创建时间
	UpdatedAt		time.Time		更新时间
	DeletedAt		gorm.DeletedAt	删除时间
	ParentID		int				父级ID
	Depth			int				分类所在层数
	Name			string			分类名
	Desc			string			分类描述
*/
type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	ParentID  int            `gorm:"column:parent_id; type: int; not null" json:"parent_id"`
	Depth     int            `gorm:"column:depth; type: int; not null" json:"depth"`
	Name      string         `gorm:"column:name; type: varchar(255);not null" json:"name"`
	Desc      string         `gorm:"column:desc; type: longtext" json:"desc"`
}
