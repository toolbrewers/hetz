package password

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string, cost int) (string, error) {
	saltedPassword, err := saltPasswordFromEnv(password)
	if err != nil {
		return "", err
	}
	return HashNoSalt(saltedPassword, cost)
}

func HashNoSalt(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func saltPasswordFromEnv(password string) (string, error) {
	salt := os.Getenv("SALT")
	if salt == "" {
		return "", errors.New("SALT environment variable not set")
	}
	return saltPassword(password, salt), nil
}

func saltPassword(password string, salt string) string {
	return password + salt
}
