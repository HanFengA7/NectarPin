package tools

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

/*
RandSalt [ 生成随机盐 ] [ 231219 ] [ 0.1 ]

------------------------------------------------------------------------------------------------------------------------

	传参: [int] values
	//: 生成盐的字节数 推荐:8,16,32,64

------------------------------------------------------------------------------------------------------------------------

	回参: [[]byte] salt [string] saltF [int] statusCode
	//: [[]byte] 生成盐的原始数据
	//: [string] 生成盐的格式化数据
	//: [statusCode] 500 --> FAIL | 200--> SUCCESS

------------------------------------------------------------------------------------------------------------------------
*/
func RandSalt(values int) (salt []byte, saltF string, statusCode int) {
	salt = make([]byte, values)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, "nil", 500
	}
	return salt, base64.StdEncoding.EncodeToString(salt), 200
}
