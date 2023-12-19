package Models

import (
	"NectarPin/constant"
	"NectarPin/tools/errmsg"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
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
	LastLonginIPAddress string         `gorm:"column:last_longin_ip_address; type: varchar(255);"json:"last_longin_ip_address,omitempty"`
	LastLonginDate      string         `gorm:"column:last_longin_date; type: varchar(255);datetime;not null" json:"last_longin_date,omitempty"`
}

/*
ExistUser [ 用户是否存在 ] [ 231211 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [int|string] values
	//: values [int] 查询用户ID | values [string] 查询用户名

------------------------------------------------------------------------------------------------------------------------

	回参: [int] msgData [int] statusCode
	//: [msgData] 0 --> 不存在 | 1 --> 存在
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func ExistUser(values interface{}) (msgData int, statusCode int) {
	var user User
	db := constant.DB
	switch values.(type) {
	case int:
		//查询用户ID
		err := db.Where("id = ?", values).First(&user).Error
		if err != nil {
			return 0, errmsg.ERROR
		}
		return 1, errmsg.SUCCESS
	case string:
		//查询用户名
		err := db.Where("username = ?", values).First(&user).Error
		if err != nil {
			return 0, errmsg.ERROR
		}
		return 1, errmsg.SUCCESS
	default:
		return 0, errmsg.ERROR
	}
}

/*
GetUser [ 查询单个用户信息 ] [ 231212 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [int] key [int|string] values
	//: key [int] --> 0 前台使用 | key [int] --> 1 后台使用
	//: values [int] 查询用户ID  | values [string] 查询用户名

------------------------------------------------------------------------------------------------------------------------

	回参:  [[]User] msgData [int] statusCode
	//: [key=>0 msgData] 用户信息
	//: (ID、Username、NickName、Email、AvatarUrl、Role)
	//: [key=>1 msgData] 用户信息
	//: (ID、CreatedAt、UpdatedAt、Username、NickName、Email、AvatarUrl、Role、LastLonginIPAddress、LastLonginDate)
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func GetUser(key int, values interface{}) (msgData []User, statusCode int) {
	var user []User
	db := constant.DB
	if key == 0 {
		switch values.(type) {
		case int:
			err := db.Select("id,username,nickname,email,avater_url,role").Where("id = ?", values).
				Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		case string:
			err := db.Select("id,username,nickname,email,avater_url,role").Where("username = ?", values).
				Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		default:
			return nil, errmsg.ERROR
		}
	} else {
		switch values.(type) {
		case int:
			err := db.Select(
				"id,created_at,updated_at,username,nickname,email,avater_url,role,last_longin_ip_address,last_longin_date",
			).Where("id = ?", values).Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		case string:
			err := db.Select(
				"id,created_at,updated_at,username,nickname,email,avater_url,role,last_longin_ip_address,last_longin_date",
			).Where("username = ?", values).Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		default:
			return nil, errmsg.ERROR
		}
	}
}

//todo 查询用户列表

/*
CreateUser [ 增加用户 ] [ 231213 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [*User] values
	//: 用户数据

------------------------------------------------------------------------------------------------------------------------

	回参: [int] msgData [int] statusCode
	//: [data] 0 --> 增加用户失败 | 1 --> 增加用户成功
	//: [code] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func CreateUser(values *User) (msgData int, statusCode int) {
	db := constant.DB
	var code int
	values.Password, code = UserPwdEnCrypto(values.Password)
	if code == 200 {
		err := db.Create(&values).Error
		if err != nil {
			return 0, errmsg.ERROR
		}
		return 1, errmsg.SUCCESS
	} else {
		return 1, errmsg.SUCCESS
	}
}

//todo 修改用户信息

/*
DeleteUser [ 删除用户 ] [ 231212 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [int|string] values
	//: values [int] 删除用户ID  | values [string] 删除用户名

------------------------------------------------------------------------------------------------------------------------

	回参: [int] msgData [int] statusCode
	//: [msgData] 0 --> 删除失败 | 1 --> 删除成功
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func DeleteUser(values interface{}) (msgData int, statusCode int) {
	var user User
	db := constant.DB
	switch values.(type) {
	case int:
		err := db.Unscoped().Delete(user, values).Error
		if err != nil {
			return 0, errmsg.ERROR
		}
		return 1, errmsg.SUCCESS
	case string:
		err := db.Where("username = ?", values).Unscoped().Delete(user).Error
		if err != nil {
			return 0, errmsg.ERROR
		}
		return 1, errmsg.SUCCESS
	default:
		return 0, errmsg.ERROR
	}
}

/*
UserPwdEnCrypto [ 用户密码加密 ] [ 231218 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参:[string] plaintext
	//:密码明文

------------------------------------------------------------------------------------------------------------------------

	回参: [string] msgData [int] statusCode
	//: [msgData] 明文密码加密后的数据
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func UserPwdEnCrypto(plaintext string) (msgData string, statusCode int) {
	salt, err := base64.StdEncoding.DecodeString(constant.Config.System.PwdHashKey)
	if err != nil {
		return "配置文件PwdHashKey解密失败", 500
	}
	ciphertext, err := scrypt.Key([]byte(plaintext), salt, 32768, 8, 2, 32)
	if err != nil {
		return "加密失败", 500
	}
	return base64.StdEncoding.EncodeToString(ciphertext), 200
}
