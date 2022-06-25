package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(password []byte, hashedpassword []byte) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
		return false
	}
	return true
}
