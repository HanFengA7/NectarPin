package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort = errors.New("密码长度不足")
	ErrPasswordHashFail = errors.New("密码哈希失败")
	ErrPasswordVerify   = errors.New("密码验证失败")
)

const (
	BcryptCost = 12
	MinPasswordLength = 6
)

func HashPassword(md5Password string) (string, error) {
	if len(md5Password) != 32 {
		return "", ErrPasswordTooShort
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(md5Password), BcryptCost)
	if err != nil {
		return "", ErrPasswordHashFail
	}

	return string(hashedBytes), nil
}

func VerifyPassword(md5Password string, hashedPassword string) bool {
	if len(md5Password) != 32 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(md5Password))
	return err == nil
}

func NeedsRehash(hashedPassword string) bool {
	hashCost, err := bcrypt.Cost([]byte(hashedPassword))
	if err != nil {
		return true
	}
	return hashCost < BcryptCost
}
