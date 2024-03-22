package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID int
	jwt.RegisteredClaims
}

func GenerateToken(userID, validHour int, secret string) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(validHour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Validate token and return the user ID
func ValidateToken(token, secret string) (int, error) {
	claims := JwtClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	return claims.UserID, nil
}
