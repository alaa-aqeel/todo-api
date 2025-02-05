package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func MakeHash(value string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyHash(value string, hashValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(value))

	return err == nil
}
