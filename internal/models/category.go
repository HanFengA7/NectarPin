package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	/*CID      int    `gorm:"column: cid; type: int; primaryKey; autoIncrement" json:"cid" label:"分类ID"`*/
	ParentID int    `gorm:"column:parent_id; type: int; not null" json:"parent_id" label:"父级分类ID"`
	LevelNum int    `gorm:"column:level_num; type: int; not null" json:"level_num" label:"分类所在层数"`
	Name     string `gorm:"column:name; type: varchar(255);not null" json:"name" label:"分类名"`
	Desc     string `gorm:"column:desc; type: longtext" json:"desc" label:"分类描述"`
}
