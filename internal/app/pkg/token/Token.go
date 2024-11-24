package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type PayloadToken struct {
	AuthId int       `json:"auth_id"`
	Exp    time.Time `json:"exp"`
}

const SecretKey = "HyVQNmB3SMjwYvL4Tqh90N7tD6ccoF8t"

func GenerateToken(tok *PayloadToken) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	tok.Exp = expirationTime

	claims := jwt.MapClaims{
		"auth_id": tok.AuthId,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokString string) (*PayloadToken, error) {
	tok, err := jwt.Parse(tokString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("unauthorized")
	}

	if claims["auth_id"] == nil {
		return nil, errors.New("unauthorized")
	}

	payloadToken := PayloadToken{
		AuthId: int(claims["auth_id"].(float64)),
		Exp:    time.Unix(int64(claims["exp"].(float64)), 0),
	}

	if time.Now().After(payloadToken.Exp) {
		return nil, errors.New("token has expired")
	}

	return &payloadToken, nil
}
