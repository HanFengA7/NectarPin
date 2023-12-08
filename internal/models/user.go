package models

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username            string `gorm:"column:username; type: varchar(255); not null" json:"username,omitempty" label:"用户名"`
	NickName            string `gorm:"column:nickname; type:varchar(255); not null" json:"nickname,omitempty" label:"昵称"`
	Password            string `gorm:"column:password; type: varchar(255); not null" json:"password,omitempty" label:"密码"`
	Email               string `gorm:"column:email; type: varchar(255); not null" json:"email,omitempty" label:"电子邮箱"`
	AvatarUrl           string `gorm:"column:avater_url; type: longtext" json:"avater_url,omitempty" label:"头像外链"`
	Role                int    `gorm:"column:role; type: int; DEFAULT:2; not null" json:"role,omitempty" validate:"required,gte=1" label:"角色码"`
	LastLonginIPAddress string `gorm:"column:last_longin_ip_address; type: varchar(255);" json:"last_longin_ip_address,omitempty" label:"最后登录IP地址"`
	LastLonginDate      string `gorm:"column:last_longin_date; type: varchar(255);datetime;not null" json:"last_longin_date,omitempty" label:"最后登录时间"`
}

//
