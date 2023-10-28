package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CompareAndVerifyPassword(hashPassword string, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), password); err != nil {
		return false
	}
	return true
}
