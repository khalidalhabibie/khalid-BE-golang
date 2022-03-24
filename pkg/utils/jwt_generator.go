package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gokes/app/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Tokens struct to describe tokens object.

// GenerateNewTokens func for generate a new Access & Refresh tokens for jamaah.
func GenerateNewTokens(id uuid.UUID, username string) (*models.Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id, username)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	// refreshToken, err := generateNewRefreshTokenJamaah()
	// if err != nil {
	// 	// Return token generation error.
	// 	return nil, err
	// }

	return &models.Tokens{
		Access: accessToken,
		// Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id uuid.UUID, username string) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY_USER")

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["username"] = username

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha512.New()

	// Create a new now date and time string with salt.
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_USER_EXPIRE_HOURS_COUNT"))

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
