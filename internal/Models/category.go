package Models

import (
	"NectarPin/constant"
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
	ShortName		string			分类缩略名
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
	ShortName string         `gorm:"column:short_name; type: varchar(255);not null" json:"short_name"`
	Desc      string         `gorm:"column:desc; type: longtext" json:"desc"`
}

// ExistCategory 检查分类是否存在[0]
func ExistCategory(name string, shortName string) (msgData string, statusCode bool) {
	var category Category
	db := constant.DB

	if db.Where("name = ?", name).Or("short_name = ?", shortName).Find(&category).RowsAffected > 0 {
		return "分类名或分类缩略名已存在", false
	}

	return "分类名或分类缩略名不存在", true
}

// ExistCategory1 检查分类是否存在[1]
func ExistCategory1(id int) (msgData string, statusCode bool) {
	var category Category
	db := constant.DB

	if db.Where("id = ?", id).Find(&category).RowsAffected > 0 {
		return "分类存在", true
	}
	return "分类不存在", false
}

// CreateCategory 创建分类
func CreateCategory(data *Category) (msgData string, statusCode int) {
	db := constant.DB

	if err := db.Create(&data).Error; err != nil {
		return "分类存入数据库失败!", 500
	}
	return "创建分类成功！", 200
}

// DeleteCategory 删除分类
func DeleteCategory(id int) (msgData string, statusCode int) {
	var category Category
	var err error
	db := constant.DB

	if id != 0 {
		err = db.Where("id = ?", id).Or("parent_id = ?", id).Unscoped().Delete(&category).Error
		if err != nil {
			return "删除分类失败!", 500
		}
		return "删除分类成功", 200
	} else {
		return "删除分类失败!", 500
	}
}

// EditCategory 编辑分类
func EditCategory(id int, oldData *Category) (msgData string, statusCode int) {
	var category Category
	var err error
	var maps = make(map[string]interface{})
	db := constant.DB

	//要编辑的数据
	maps["name"] = oldData.Name
	maps["short_name"] = oldData.ShortName
	maps["desc"] = oldData.Desc
	maps["parent_id"] = oldData.ParentID
	maps["depth"] = oldData.Depth

	//判断分类是否存在
	ExistCategoryMsg, ExistCategoryBool := ExistCategory1(id)
	if ExistCategoryBool != true {
		return ExistCategoryMsg, 500
	}

	//编辑分类信息
	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return "编辑分类信息失败", 500
	}
	return "编辑分类信息成功", 200
}

// GetCategory 查询分类
func GetCategory(id int) (msgData []Category, statusCode int) {
	var category []Category
	db := constant.DB

	err := db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, 500
	}

	return category, 200
}

// GetCategoryList 查询分类列表
func GetCategoryList(pageSize int, pageNum int) (data []Category, total int64, statusCode int) {

	if pageSize <= 0 {
		//每页的数据条数
		pageSize = 5
	}
	if pageNum <= 0 {
		//当前请求的页数
		pageNum = 1
	}

	var err error
	db := constant.DB
	offset := (pageNum - 1) * pageSize

	err = db.Limit(pageSize).Offset(offset).Find(&data).Error
	err = db.Model(&data).Count(&total).Error
	if err != nil {
		return nil, 0, 500
	}
	return data, total, 200
}
