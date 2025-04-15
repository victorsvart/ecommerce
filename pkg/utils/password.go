package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashed, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(plainPassword *string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	*plainPassword = string(hashed)
	return nil
}
