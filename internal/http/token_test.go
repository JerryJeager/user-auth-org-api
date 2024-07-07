package http

import (
	"net/http"
	"os"
	"testing"

	"github.com/JerryJeager/user-auth-org-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	// Set up
	secret := "test_secret"
	os.Setenv("API_SECRET", secret)
	defer os.Unsetenv("API_SECRET")

	userID := uuid.New()

	tokenString, err := utils.GenerateToken(userID)
	assert.NoError(t, err)

	c, _ := gin.CreateTestContext(nil)
	c.Request = &http.Request{
		Header: http.Header{
			"Authorization": []string{"Bearer " + tokenString},
		},
	}

	id, err := ValidateToken(c)

	assert.NoError(t, err)
	assert.Equal(t, userID.String(), id)
}
