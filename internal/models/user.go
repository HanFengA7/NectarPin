package models

import (
	"gorm.io/gorm"
	"time"
)

/*
User [用户表结构][231209][0.1]

	ID						uint				用户ID
	CreatedAt				time.Time			创建时间
	UpdatedAt				time.Time			更新时间
	DeletedAt				gorm.DeletedAt		删除时间
	Username            	string				用户名
	NickName            	string				昵称
	Password            	string				密码
	Email               	string				电子邮箱
	AvatarUrl          		string				头像地址
	Role                	int					角色码 [0:管理员 1:编辑者 2:订阅者]
	LastLonginIPAddress 	string				最后登录IP地址
	LastLonginDate      	string				最后登录时间
*/
type User struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time      `json:"created_at,omitempty"`
	UpdatedAt           time.Time      `json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Username            string         `gorm:"column:username; type: varchar(255); not null" json:"username,omitempty"`
	NickName            string         `gorm:"column:nickname; type:varchar(255); not null" json:"nickname,omitempty"`
	Password            string         `gorm:"column:password; type: varchar(255); not null" json:"password,omitempty"`
	Email               string         `gorm:"column:email; type: varchar(255); not null" json:"email,omitempty"`
	AvatarUrl           string         `gorm:"column:avater_url; type: longtext" json:"avater_url,omitempty"`
	Role                int            `gorm:"column:role; type: int; DEFAULT:2; not null" json:"role,omitempty"`
	LastLonginIPAddress string         `gorm:"column:last_longin_ip_address; type: varchar(255);" json:"last_longin_ip_address,omitempty"`
	LastLonginDate      string         `gorm:"column:last_longin_date; type: varchar(255);datetime;not null" json:"last_longin_date,omitempty"`
}
