package token

import (
	"github.com/stretchr/testify/assert"
	"mncTest/internal/app/pkg/token"
	"mncTest/internal/utils"
	"testing"
	"time"
)

func TestAuthUtil_ValidToken(t *testing.T) {
	payload := &token.PayloadToken{
		AuthId: 123,
		Exp:    time.Now().Add(10 * time.Minute),
	}
	tokenString, err := token.GenerateToken(payload)
	assert.NoError(t, err, "Failed to generate token")
	authHeader := "Bearer " + tokenString
	customerID, err := utils.AuthUtil(authHeader)
	assert.NoError(t, err, "AuthUtil returned an unexpected error")
	assert.Equal(t, 123, customerID, "Auth ID does not match")
}

func TestAuthUtil_InvalidToken(t *testing.T) {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoxMjMsImV4cCI6MTYzMDM1NTU1Mn0.invalidsignature"
	authHeader := "Bearer " + invalidToken
	customerID, err := utils.AuthUtil(authHeader)
	assert.Error(t, err, "Expected error due to invalid token")
	assert.Equal(t, 0, customerID, "Expected customerID to be 0 due to invalid token")
}

func TestAuthUtil_InvalidHeader(t *testing.T) {
	authHeader := "Bearer "
	customerID, err := utils.AuthUtil(authHeader)
	assert.Error(t, err, "Expected error due to invalid authorization header")
	assert.Equal(t, 0, customerID, "Expected customerID to be 0 due to invalid header")
}

func TestAuthUtil_EmptyHeader(t *testing.T) {
	authHeader := ""
	customerID, err := utils.AuthUtil(authHeader)
	assert.Error(t, err, "Expected error due to empty authorization header")
	assert.Equal(t, 0, customerID, "Expected customerID to be 0 due to empty header")
}
