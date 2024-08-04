package tokens

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string
	UID      string
	jwt.RegisteredClaims
}

func GenerateJwt(username, uid string) (string, error) {
	return "", nil
}

func VerifyJwt() (string, error) {
	return "", nil
}
