package util

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
)

func ScryptPassword(password string) (string, error) {
	salt := make([]byte, 8)
	salt = []byte{23, 54, 69, 21, 78, 37, 11, 5}
	if hash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32); err == nil {
		return base64.StdEncoding.EncodeToString(hash), nil
	} else {
		return "", err
	}
}
