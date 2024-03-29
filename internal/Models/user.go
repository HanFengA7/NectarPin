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
	PSignatures				string				个性签名
	Role                	int					角色码 [0:管理员 1:编辑者 2:订阅者]
	ThisLonginIPAddress 	string				这次登录IP地址
	ThisLonginDate      	string				这次登录时间
	LastLonginIPAddress 	string				最后登录IP地址
	LastLonginDate      	string				最后登录时间
*/
type User struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time      `json:"created_at,omitempty"`
	UpdatedAt           time.Time      `json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Username            string         `gorm:"column:username; type: varchar(255); not null" json:"username"`
	NickName            string         `gorm:"column:nickname; type:varchar(255); not null" json:"nickname"`
	Password            string         `gorm:"column:password; type: varchar(255); not null" json:"password,omitempty"`
	Email               string         `gorm:"column:email; type: varchar(255); not null" json:"email"`
	AvatarUrl           string         `gorm:"column:avater_url; type: longtext" json:"avater_url"`
	PSignatures         string         `gorm:"column:p_signatures; type: longtext" json:"p_signatures"`
	Role                int            `gorm:"column:role; type: int; DEFAULT:2; not null" json:"role"`
	ThisLonginIPAddress string         `gorm:"column:this_longin_ip_address; type: varchar(255);" json:"this_longin_ip_address,omitempty"`
	ThisLonginDate      string         `gorm:"column:this_longin_date; type: varchar(255);datetime;not null" json:"this_longin_date"`
	LastLonginIPAddress string         `gorm:"column:last_longin_ip_address; type: varchar(255);" json:"last_longin_ip_address,omitempty"`
	LastLonginDate      string         `gorm:"column:last_longin_date; type: varchar(255);datetime;not null" json:"last_longin_date"`
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
			err := db.Select("id,username,nickname,email,avater_url,role,p_signatures").Where("id = ?", values).
				Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		case string:
			err := db.Select("id,username,nickname,email,avater_url,role,p_signatures").Where("username = ?", values).
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
				"id,created_at,updated_at,username,nickname,email,avater_url,role,p_signatures,last_longin_ip_address,last_longin_date,this_longin_ip_address,this_longin_date",
			).Where("id = ?", values).Find(&user).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errmsg.ERROR
			}
			return user, errmsg.SUCCESS
		case string:
			err := db.Select(
				"id,created_at,updated_at,username,nickname,email,avater_url,role,p_signatures,last_longin_ip_address,last_longin_date,this_longin_ip_address,this_longin_date",
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

/*
GetUserList [ 查询用户列表 ] [ 231220 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: ([string] username 可选) [int] pageSize [int] pageNum
	//: [username] 模糊搜索用户名类似的用户数据列表,是可选项
	//: [pageSize] 代表每页的数据条数,是必选项
	//: [pageNum]  代表当前请求的页数,是必选项

------------------------------------------------------------------------------------------------------------------------

	回参: [[]User] data [int] total [int] statusCode
	//: [data] 列表数据
	//: [total] 数据总量
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func GetUserList(username string, pageSize int, pageNum int) (data []User, total int64, statusCode int) {
	if pageSize <= 0 {
		pageSize = 5
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	var err error
	db := constant.DB
	offset := (pageNum - 1) * pageSize

	switch {
	case len(username) == 0:
		err = db.Select(
			"id,created_at,updated_at,username,nickname,email,avater_url,role,p_signatures,last_longin_ip_address,last_longin_date,this_longin_ip_address,this_longin_date",
		).Limit(pageSize).Offset(offset).Find(&data).Error
		err = db.Model(&data).Count(&total).Error
		if err != nil {
			return nil, 0, 500
		}
		return data, total, 200

	case len(username) != 0:
		err = db.Select(
			"id,created_at,updated_at,username,nickname,email,avater_url,role,p_signatures,last_longin_ip_address,last_longin_date,this_longin_ip_address,this_longin_date",
		).Where("username LIKE ?", username+"%").Limit(pageSize).Offset(offset).Find(&data).Error
		err = db.Model(&data).Where("username LIKE ?", username+"%").Count(&total).Error
		if err != nil {
			return nil, 0, 500
		}
		return data, total, 200

	default:
		return nil, 0, 500
	}

}

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

/*
EditUserInfo [ 修改用户信息 ] [ 231220 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [int] id [string] values
	//: [id] 用户的ID
	//: [values] 修改的用户信息

------------------------------------------------------------------------------------------------------------------------

	回参: [string] msgData [int] statusCode
	//: [msgDate] 返回的信息
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/

func EditUserInfo(id int, values *User) (msgData string, statusCode int) {
	var err error
	var checkCodeA int
	var user User
	var maps = make(map[string]interface{})
	db := constant.DB

	maps["username"] = values.Username
	maps["nickname"] = values.NickName
	maps["email"] = values.Email
	maps["avater_url"] = values.AvatarUrl
	maps["p_signatures"] = values.PSignatures
	//maps["role"] = values.Role

	//判断用户名是否重复
	rows := db.Select("id,username").Where("username = ?", values.Username).First(&user).RowsAffected
	switch {
	case rows == 0:
		checkCodeA = 1
	case rows > 0:
		if user.ID == uint(id) {
			checkCodeA = 1
		} else {
			checkCodeA = 0
		}
	}

	switch {
	case checkCodeA == 1:
		err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
		if err != nil {
			return "编辑用户信息失败", 500
		}
		return "编辑用户信息成功", 200
	default:
		return "username已存在,无法修改", 500
	}
}

/*
EditUserPwd [ 修改用户密码 ] [ 240208 ] [ 0.2 ]

------------------------------------------------------------------------------------------------------------------------

	传参:[string] username [string] password
	//: [username] 用户名
	//: [old_password] 旧密码(在前端进行了MD5加密)
	//: [new_password] 新密码(在前端进行了MD5加密)

------------------------------------------------------------------------------------------------------------------------

	回参: [int] msgData [int] statusCode
	//: [msgData] 修改密码失败 | 修改密码成功
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func EditUserPwd(username string, oldPassword string, newPassword string) (msgData string, statusCode int) {
	db := constant.DB
	var err error
	var user User

	oldPassword, _ = UserPwdEnCrypto(oldPassword)
	err = db.Where("username = ?", username).Where("password = ?", oldPassword).First(&user).Error
	if err != nil {
		return "旧密码错误", 500
	}

	err = db.Model(&user).Where("username = ?", username).Update("password", newPassword).Error
	if err != nil {
		return "编辑用户密码失败", 500
	}
	return "编辑用户密码成功", 200
}

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

/*
CheckLogin [ 用户登录验证 ] [ 240115 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参:[string] username [string] plaintext
	//: [username] 用户名
	//: [plaintext] 密码明文(在前端进行了MD5加密)

------------------------------------------------------------------------------------------------------------------------

	回参: [string] msgData [int] statusCode
	//: [msgData] 验证成功 或 验证失败,请检查用户名或密码!
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func CheckLogin(username string, plaintext string) (msgData string, startCode int) {
	var db = constant.DB
	var err error
	var user User

	err = db.Where("username = ?", username).First(&user).Error
	//passwordMd5 := fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	jmPwd, _ := UserPwdEnCrypto(plaintext)

	if err != nil {
		return "验证失败,请检查用户名或密码!", 500
	}
	if user.ID == 0 {
		return "验证失败,请检查用户名或密码!", 500
	}
	if user.Password == jmPwd {
		return "验证成功", 200
	} else {
		return "验证失败,请检查用户名或密码!", 500
	}
}

/*
UserLoginWriteLog [ 写入用户登录日志 ] [ 240201 ] [ 0.1 ]
------------------------------------------------------------------------------------------------------------------------

	//:[入参]
	[1] [uID] [int] : 用户ID
	[2] [LIP] [string] : 登录时的IP地址
	[3] [lDate] [string] : 登录时的时间

------------------------------------------------------------------------------------------------------------------------
*/
func UserLoginWriteLog(uID int, lIP string, lDate string) {
	var db = constant.DB
	var user User
	var maps = make(map[string]interface{})
	db.Where("id = ?", uID).First(&user)
	maps["ThisLonginIPAddress"] = lIP
	maps["ThisLonginDate"] = lDate
	maps["LastLonginIPAddress"] = user.ThisLonginIPAddress
	maps["LastLonginDate"] = user.ThisLonginDate
	db.Model(&user).Where("id = ?", uID).Select("LastLonginIPAddress", "LastLonginDate", "ThisLonginIPAddress", "ThisLonginDate").Updates(&maps)
}
