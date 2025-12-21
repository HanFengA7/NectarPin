package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT 声明结构
type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

/*
GenerateToken 生成 JWT Token
@param userID uint 用户ID
@param userName string 用户名
@param role string 用户角色
@param secret string JWT 密钥
@param expireHours int 过期时间（小时）
@return token string 生成的 Token
@return err error 错误信息
*/
func GenerateToken(userID uint, userName string, role string, secret string, expireHours int) (string, error) {
	// 创建声明
	claims := Claims{
		UserID:   userID,
		UserName: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "nectarpin",
		},
	}

	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*
ParseToken 解析 JWT Token
@param tokenString string Token 字符串
@param secret string JWT 密钥
@return claims *Claims 声明信息
@return err error 错误信息
*/
func ParseToken(tokenString string, secret string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

/*
RefreshToken 刷新 Token
@param tokenString string 原 Token 字符串
@param secret string JWT 密钥
@param expireHours int 新的过期时间（小时）
@return newToken string 新的 Token
@return err error 错误信息
*/
func RefreshToken(tokenString string, secret string, expireHours int) (string, error) {
	// 解析旧 token
	claims, err := ParseToken(tokenString, secret)
	if err != nil {
		return "", err
	}

	// 生成新 token
	return GenerateToken(claims.UserID, claims.UserName, claims.Role, secret, expireHours)
}
