package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hash)
}

func ComparePassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
