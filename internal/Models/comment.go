package Models

import (
	"gorm.io/gorm"
	"time"
)

/*
Comment [评论表结构][231209][0.1]

	ID				uint			评论ID
	CreatedAt		time.Time		创建时间
	UpdatedAt		time.Time		更新时间
	DeletedAt		gorm.DeletedAt	删除时间
	AID   			int				文章ID
	ReplyID			int				回复ID
	RootID			int				评论根ID
	Depth			int				当前评论层
	CUID			string			评论用户ID(时间MD5)
	CUsername		string			评论用户名
	CEmail			string			评论用户邮箱
	CWebSite		string			评论用户网站
	CIP 			string			评论IP
	CContent 		string			评论内容
	CFlag			int				评论者标志 [0:管理员 1:游客]
	CLevel			int				评论等级 [0:正常 1:精选评论 2:仅管理员可见 3:隐藏]
	CUpvote		    int				评论点赞数
*/
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Article   Article        `gorm:"foreignKey:AID"`
	AID       int            `gorm:"column:aid; not null" json:"aid,omitempty"`
	ReplyID   int            `gorm:"column:reply_id" json:"reply_id,omitempty"`
	RootID    int            `gorm:"column:root_id; not null" json:"root_id,omitempty"`
	Depth     int            `gorm:"column:depth; not null" json:"depth,omitempty"`
	CUID      string         `gorm:"column:cuid; not null" json:"cuid,omitempty"`
	CUsername string         `gorm:"column:c_username; not null" json:"c_username,omitempty"`
	CEmail    string         `gorm:"column:c_email; not null" json:"c_email,omitempty"`
	CWebSite  string         `gorm:"column:c_website" json:"c_website,omitempty"`
	CIP       string         `gorm:"column:cip; not null" json:"cip,omitempty"`
	CContent  string         `gorm:"column:c_content; not null; type:longtext" json:"c_content,omitempty"`
	CFlags    int            `gorm:"column:c_flags; not null" json:"c_flags,omitempty"`
	CLevel    int            `gorm:"column:c_level; not null" json:"c_level,omitempty"`
	CUpvote   int            `gorm:"column:c_upvote" json:"c_upvote,omitempty"`
}
