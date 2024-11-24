package utils

import (
	"errors"
	"mncTest/internal/app/pkg/token"
	"strings"
)

func AuthUtil(authHeader string) (customerID int, err error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return 0, errors.New("invalid authorization header")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	payloadToken, err := token.ValidateToken(tokenString)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	return payloadToken.AuthId, nil
}
