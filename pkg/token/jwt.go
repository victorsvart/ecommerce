package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint64, issuerEM string) (string, error) {
	claims := AppClaims{
		jwt.RegisteredClaims{
			ID:        string(userID),
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

func ParseJWT(tokenString string) (AppClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return AppClaims{}, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		secret := os.Getenv("APP_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		return AppClaims{}, err
	}

	if claims, ok := token.Claims.(AppClaims); ok && token.Valid {
		return claims, nil
	}

	return AppClaims{}, fmt.Errorf("invalid token")
}
