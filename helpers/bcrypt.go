package helpers

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(plain string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)

}
func PassCompare(hash string, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}
	return true
}
func Base64(data string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}
