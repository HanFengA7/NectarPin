package middleware

import (
	"NectarPin/constant"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

/*
CreateToken [创建Token] [240117] [0.1]
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
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(JwtHashKey)
	if err != nil {
		return "", 500, "创建Token失败"
	}
	return token, 200, "Token创建完成"
}
