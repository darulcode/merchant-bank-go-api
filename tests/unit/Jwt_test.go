package token

import (
	"mncTest/internal/app/pkg/token"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	payload := &token.PayloadToken{
		AuthId: "123",
		Exp:    time.Now().Add(10 * time.Minute),
	}
	tokenString, err := token.GenerateToken(payload)

	assert.NoError(t, err, "GenerateToken returned an unexpected error")
	assert.NotEmpty(t, tokenString, "Generated token should not be empty")
	assert.WithinDuration(t, payload.Exp, time.Now().Add(10*time.Minute), time.Second, "Expiration time should be approximately 10 minutes from now")
}

func TestValidateToken(t *testing.T) {
	payload := &token.PayloadToken{
		AuthId: "123",
		Exp:    time.Now().Add(10 * time.Minute),
	}
	tokenString, err := token.GenerateToken(payload)
	assert.NoError(t, err, "GenerateToken returned an unexpected error")
	validPayload, err := token.ValidateToken(tokenString)

	assert.NoError(t, err, "ValidateToken returned an unexpected error")
	assert.NotNil(t, validPayload, "ValidateToken should return a non-nil payload")
	assert.Equal(t, payload.AuthId, validPayload.AuthId, "AuthId in validated token should match the original")
	assert.WithinDuration(t, payload.Exp, validPayload.Exp, time.Second, "Expiration time in validated token should match the original")
}

func TestValidateExpiredToken(t *testing.T) {
	payload := &token.PayloadToken{
		AuthId: "123",
		Exp:    time.Now().Add(1 * time.Second),
	}
	tokenString, err := token.GenerateToken(payload)
	assert.NoError(t, err, "GenerateToken returned an unexpected error")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	time.Sleep(2 * time.Second)
	_, err = token.ValidateToken(tokenString)
	assert.EqualError(t, err, "token has expired", "Error message should indicate token expiration")
}

func TestValidateInvalidToken(t *testing.T) {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoxMjMsImV4cCI6MTYzMDM1NTU1Mn0.invalidsignature"
	_, err := token.ValidateToken(invalidToken)

	assert.Error(t, err, "ValidateToken should return an error for an invalid token")
	assert.EqualError(t, err, "signature is invalid", "Error message should indicate unauthorized access")
}
