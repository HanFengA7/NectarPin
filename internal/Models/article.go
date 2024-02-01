package Models

import (
	"NectarPin/constant"
	"gorm.io/gorm"
	"time"
)

/*
Article [文章表结构][231209][0.1]

	ID						uint				文章ID
	CreatedAt				time.Time			创建时间
	UpdatedAt				time.Time			更新时间
	DeletedAt				gorm.DeletedAt		删除时间
	UID           			int					用户ID(外键)
	CID           			int					分类ID(外键)
	Title         			string				文章标题
	Tags          			string				文章标签
	Desc          			string				文章描述
	Content       			string				文章内容
	ImgUrl        			string				文章展图
	AIFComment    			int					是否开启评论[0:关闭 1:开启]
	AIFHide       			int					是否隐藏文章[0:关闭 1:开启]
	AIFEncrypt    			int					是否加密文章[0:关闭 1:开启]
	AIFEncryptPwd 			string				加密文章的密码
*/
type Article struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	User          User           `gorm:"foreignKey:UID"`
	Category      Category       `gorm:"foreignKey:CID"`
	UID           int            `gorm:"column:uid; type:int; not null" json:"uid" `
	CID           int            `gorm:"column:cid; type:int; not null" json:"cid" `
	Title         string         `gorm:"column:title; type:varchar(255);not null" json:"title" label:"文章标题" `
	Tags          string         `gorm:"column:tags; type:text" json:"tags" label:"文章标签"`
	Desc          string         `gorm:"column:desc; type:varchar(255)" json:"desc" label:"文章概述"`
	Content       string         `gorm:"column:content; type:longtext;not null" json:"content"  label:"文章正文" `
	ImgUrl        string         `gorm:"column:img_url; type:longtext" json:"img_url" label:"文章展图"`
	AIFComment    int            `gorm:"column:aif_comment; type:int; not null;" json:"aif_comment" label:"是否开启评论"`
	AIFHide       int            `gorm:"column:aif_hide; type:int; not null;" json:"aif_hide" label:"是否隐藏文章"`
	AIFEncrypt    int            `gorm:"column:aif_encrypt; type:int; not null;" json:"aif_encrypt" label:"是否加密文章"`
	AIFEncryptPwd string         `gorm:"column:aif_encrypt_pwd; type:string;" json:"aif_encrypt_pwd,-" label:"加密文章密码"`
}

/*
CreateArticle 创建文章 [0.1] [240130]
------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1] [data] [*Article] : 文章的数据

------------------------------------------------------------------------------------------------------------------------

	//:[出参]
	[1] [msgData] [string] : 返回的消息
	[2] [statusCode] [int] : 返回的状态码

------------------------------------------------------------------------------------------------------------------------
*/
func CreateArticle(data *Article) (msgData string, statusCode int) {
	db := constant.DB
	if err := db.Create(&data).Error; err != nil {
		return "文章存入数据库失败!", 500
	}
	return "创建文章成功！", 200
}

/*
GetArticle 查询文章 [0.1] [240131]
------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1] [id] [int] : 文章的ID

------------------------------------------------------------------------------------------------------------------------

	//:[出参]
	[1] [msgData] [[]Article] : 返回的数据
	[2] [statusCode] [int] : 返回的状态码

------------------------------------------------------------------------------------------------------------------------
*/
func GetArticle(id int) (msgData []Article, statusCode int) {
	var article []Article
	db := constant.DB
	if err := db.Joins(
		"User", db.Select("id", "username", "nickname", "email", "avater_url", "role", "last_longin_date"),
	).Joins("Category").Where("article.id = ?", id).First(&article).Error; err != nil {
		return nil, 500
	}
	return article, 200
}
