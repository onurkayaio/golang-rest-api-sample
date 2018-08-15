package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}
