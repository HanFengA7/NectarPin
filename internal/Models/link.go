package Models

import (
	"gorm.io/gorm"
	"time"
)

/*
Link [友情链接表结构][231209][0.1]

	ID						uint				友链ID
	CreatedAt				time.Time			创建时间
	UpdatedAt				time.Time			更新时间
	DeletedAt				gorm.DeletedAt		删除时间
	Name      				string				名称
	Url       				string				链接
	Logo      				string				Logo
	Desc     				string				描述
	Category 				string				分类
*/
type Link struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name      string         `gorm:"column:name; not null" json:"name"`
	Url       string         `gorm:"column:url; not null" json:"url"`
	Logo      string         `gorm:"column:logo; not null" json:"logo"`
	Desc      string         `gorm:"column:desc; not null" json:"desc"`
	Category  string         `gorm:"column:category" json:"category"`
}
