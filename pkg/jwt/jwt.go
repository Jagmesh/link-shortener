package jwt

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

type claimsKeyType string

const (
	CLAIMS_CTX_KEY claimsKeyType = "claimsCtxKey"
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

func (j *JWT) Parse(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method != j.signingMethod {
			return nil, fmt.Errorf("unexpected signing method: %v. Must be %v", t.Header["alg"], j.signingMethod)
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	return claims, nil
}

func GetClaimsFromContext[T any](ctx context.Context) (*T, error) {
	rawValue := ctx.Value(CLAIMS_CTX_KEY)
	if rawValue == nil {
		return nil, fmt.Errorf("key %v not found in context", CLAIMS_CTX_KEY)
	}

	if typedValue, ok := rawValue.(*T); ok {
		return typedValue, nil
	}

	if claims, ok := rawValue.(jwt.MapClaims); ok {
		var result T
		if err := mapstructure.Decode(claims, &result); err != nil {
			return nil, fmt.Errorf("failed to decode claims into %T: %v", result, err)
		}
		return &result, nil
	}

	return nil, fmt.Errorf("value has unexpected type %T, expected %T", rawValue, new(T))
}
