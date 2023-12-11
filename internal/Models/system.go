package Models

import (
	"gorm.io/gorm"
	"time"
)

/*
System [系统表结构][231210][0.1]

	ID						uint				文章ID
	CreatedAt				time.Time			创建时间
	UpdatedAt				time.Time			更新时间
	DeletedAt				gorm.DeletedAt		删除时间
	Key						string				Key
	Values					string				Values
	ValuesPlus				string				ValuesPlus
*/
type System struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Key        string         `gorm:"column:key; not null" json:"key,omitempty"`
	Values     string         `gorm:"column:values; not null" json:"values,omitempty"`
	ValuesPlus string         `gorm:"column:values_plus; not null " json:"values_plus,omitempty"`
}
