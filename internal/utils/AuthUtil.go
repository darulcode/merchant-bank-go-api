package utils

import (
	"errors"
	"mncTest/internal/app/pkg/token"
	"strings"
)

func AuthUtil(authHeader string) (customerID string, err error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid authorization header")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	payloadToken, err := token.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return payloadToken.AuthId, nil
}
