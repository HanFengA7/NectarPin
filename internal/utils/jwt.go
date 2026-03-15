package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("无效的令牌")
	ErrExpiredToken = errors.New("令牌已过期")
	ErrMissingToken = errors.New("缺少令牌")
)

type JWTClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey     string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

var JWTConfigInstance *JWTConfig

func InitJWT(secret string) {
	JWTConfigInstance = &JWTConfig{
		SecretKey:     secret,
		AccessExpiry:  time.Hour,
		RefreshExpiry: 7 * 24 * time.Hour,
	}
}

func GenerateAccessToken(userID uint64) (string, time.Time, error) {
	return GenerateToken(userID, JWTConfigInstance.AccessExpiry)
}

func GenerateRefreshToken(userID uint64) (string, time.Time, error) {
	return GenerateToken(userID, JWTConfigInstance.RefreshExpiry)
}

func GenerateToken(userID uint64, expiry time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(expiry)

	claims := &JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTConfigInstance.SecretKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(JWTConfigInstance.SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func ValidateToken(tokenString string) (uint64, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
