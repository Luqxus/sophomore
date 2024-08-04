package tokens

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string
	UID      string
	jwt.RegisteredClaims
}

func GenerateJwt(username, uid string) (string, error) {
	return "", nil
}

func VerifyJwt(signedToken string) (string, error) {
	claims := new(Claims)

	token, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		key := os.Getenv("SECRET_KEY")
		if key == "" {
			log.Panic("secret key not found")
			return nil, errors.New("critical server error occured")
		}

		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid authorization header")
	}

	return claims.UID, nil
}
