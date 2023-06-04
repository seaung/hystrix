package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(src string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
