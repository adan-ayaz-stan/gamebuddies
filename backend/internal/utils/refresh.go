package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const RefreshTokenExpiry = 7 * 24 * time.Hour

// GenerateRefreshToken creates a refresh token that embeds the user ID for O(1) DB lookup.
// Token format: "<userID>:<base64secret>"
func GenerateRefreshToken(userID uint) (plainToken string, hashedToken string, expiry time.Time, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", time.Time{}, err
	}

	secret := base64.URLEncoding.EncodeToString(b)
	plainToken = fmt.Sprintf("%d:%s", userID, secret)

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plainToken), 12)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return plainToken, string(hashBytes), time.Now().Add(RefreshTokenExpiry), nil
}

// ParseUserIDFromRefreshToken extracts the user ID embedded in the token.
func ParseUserIDFromRefreshToken(token string) (uint, error) {
	parts := strings.SplitN(token, ":", 2)
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid refresh token format")
	}
	id, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid refresh token format")
	}
	return uint(id), nil
}
