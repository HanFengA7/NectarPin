package tools

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func RandSalt(values int) (salt []byte, saltF string, statusCode int) {
	salt = make([]byte, values)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, "nil", 500
	}
	return salt, base64.StdEncoding.EncodeToString(salt), 200
}
