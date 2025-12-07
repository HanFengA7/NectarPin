package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
EncryptUserPassword 加密用户密码
@param password string 用户密码
@return msg string 提示信息
@return code int 状态码
@return data any 返回数据
*/
func EncryptUserPassword(password string) (msg string, code int, data string) {
	originPassword := []byte(password)
	encryptPassword, err := bcrypt.GenerateFromPassword(originPassword, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Sprintf("加密用户密码失败: %s", err.Error()), 500, ""
	}
	return "加密用户密码成功", 200, string(encryptPassword)
}

/*
CheckUserPassword 校验用户密码
@param originPassword string 加密前的密码
@param encryptPassword string 加密后的密码
@return msg string 提示信息
@return code int 状态码
@return data any 返回数据
*/
func CheckUserPassword(originPassword string, encryptPassword string) (msg string, code int) {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(originPassword))
	if err != nil {
		return "用户密码验证失败", 500
	}
	return "用户密码验证成功", 200
}
