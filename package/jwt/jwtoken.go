package jwtoken

import "github.com/golang-jwt/jwt"

type JWToken struct {
	Secret string
}

func NewJWToken(secret string) *JWToken {
	return &JWToken{Secret: secret}
}

func GenerateToken(claims *jwt.Token, secret string) (string, error) {
	hmacSecret := []byte(secret)
	token, err := claims.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
