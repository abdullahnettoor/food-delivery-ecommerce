package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(p)
	return hashedPassword, err
}

func CompareHashedPassword(dbPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, err
}
