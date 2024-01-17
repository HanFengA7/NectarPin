package middleware

import (
	"NectarPin/constant"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

/*
CreateToken [创建Token] [240117] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[入参][1][username][string]: 用户名
	[入参][2][id][int]: 用户ID
	[入参][3][role][int]: 用户角色码

------------------------------------------------------------------------------------------------------------------------

	[回参][1][token][string]: 返回的Token
	[回参][2][statusCode][int]: 返回的状态码
	[回参][3][statusMsg][string]: 返回的信息

------------------------------------------------------------------------------------------------------------------------
*/
func CreateToken(username string, id int, role int) (token string, statusCode int, statusMsg string) {
	JwtHashKey := []byte(constant.Config.System.JwtHashKey)

	type MyCustomClaims struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Role     int    `json:"role"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		id,
		username,
		role,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "NectarPin",
			ID:        strconv.Itoa(id),
			Audience:  []string{"somebody_else"},
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(JwtHashKey)
	if err != nil {
		fmt.Println(err)
		return "", 500, "创建Token失败"
	}
	return token, 200, "Token创建完成"
}
