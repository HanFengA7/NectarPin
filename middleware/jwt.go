package middleware

import (
	"NectarPin/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MyCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

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

	claims := MyCustomClaims{
		id,
		username,
		role,
		jwt.RegisteredClaims{
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

/*
VerifyToken [验证Token] [240117] [0.1]
------------------------------------------------------------------------------------------------------------------------

	[入参][1][tokenString][string]: Token

------------------------------------------------------------------------------------------------------------------------

	[回参][1][tokenData][string]: 返回Token解密后的数据
	[回参][2][tokenBool][bool]: 返回Token验证是否通过 false-->不通过 | true-->通过
	[回参][3][statusCode][int]: 返回的状态码

------------------------------------------------------------------------------------------------------------------------
*/
func VerifyToken(tokenString string) (tokenData *MyCustomClaims, tokenBool bool, statusCode int) {
	JwtHashKey := []byte(constant.Config.System.JwtHashKey)
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtHashKey, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		fmt.Println(err)
		return nil, false, 500
	} else if claims, valid := token.Claims.(*MyCustomClaims); token.Valid {
		return claims, valid, 200
	} else {
		return claims, valid, 500
	}
}

/*
AuthJWT [JWT中间件] [240118] [0.1]
*/
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		JWTHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(JWTHeader) != 0 && JWTHeader[0] == "Bearer" {
			_, VJWTBool, _ := VerifyToken(JWTHeader[1])
			if VJWTBool != false {
				c.Next()
			} else {
				c.JSON(
					http.StatusForbidden,
					gin.H{
						"code": http.StatusForbidden,
						"msg":  "身份认证失败!",
					})
				c.Abort()
			}
		} else {
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"code": http.StatusForbidden,
					"msg":  "身份认证失败!",
				})
			c.Abort()
		}
	}
}
