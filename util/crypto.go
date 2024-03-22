package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(plaintext string, hashed string) error {
	log.Println(plaintext, hashed)
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext))
}
