package utils

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// Set up
	secret := "test_secret"
	os.Setenv("API_SECRET", secret)
	defer os.Unsetenv("API_SECRET")

	userID := uuid.New()

	// Act
	tokenString, err := GenerateToken(userID)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	// Verify claims
	assert.Equal(t, true, claims["authorized"])
	assert.Equal(t, userID.String(), claims["id"])
	exp := int64(claims["exp"].(float64))
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(exp, 0), time.Minute)
}
