package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret        string
	signingMethod jwt.SigningMethod
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret:        secret,
		signingMethod: jwt.SigningMethodHS256,
	}
}

func (j *JWT) Create(claimMap jwt.MapClaims) (string, error) {
	signedToken, err := jwt.NewWithClaims(j.signingMethod, claimMap).SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JWT) Parse(token string) (any, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method != j.signingMethod {
			return nil, fmt.Errorf("unexpected signing method: %v. Must be %v", t.Header["alg"], j.signingMethod)
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken.Claims, nil
}
