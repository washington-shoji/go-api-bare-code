package helpers

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(encPass string, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(pass)) == nil
}
