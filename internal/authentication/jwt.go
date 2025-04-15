package authentication

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	uint64
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint64, issuerEM string) (string, error) {
	claims := AppClaims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    issuerEM,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", err
	}

	return ss, nil
}
