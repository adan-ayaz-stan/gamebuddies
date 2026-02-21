package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func jwtSecret() ([]byte, error) {
	s := []byte(os.Getenv("JWT_SECRET"))
	if len(s) == 0 {
		return nil, fmt.Errorf("JWT_SECRET environment variable not set")
	}
	return s, nil
}

func GenerateAccessToken(userID uint) (string, time.Time, error) {
	secret, err := jwtSecret()
	if err != nil {
		return "", time.Time{}, err
	}

	expiry := time.Now().Add(15 * time.Minute)

	claims := &AccessTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			Issuer:    "gamebuddies",
			Subject:   fmt.Sprintf("%d", userID),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiry, nil
}
