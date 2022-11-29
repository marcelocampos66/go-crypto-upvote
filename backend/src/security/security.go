package security

import "golang.org/x/crypto/bcrypt"

func Hash(text string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
