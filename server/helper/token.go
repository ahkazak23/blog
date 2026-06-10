package helper

import (
	"blog/model"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email  string
	UserId uint

	jwt.RegisteredClaims
}

func jwtSecret() ([]byte, error) {
	secret := strings.TrimSpace(os.Getenv("jwt_secret"))
	if secret == "" {
		return nil, errors.New("jwt_secret environment variable is required")
	}

	return []byte(secret), nil
}

func GenerateToken(user model.User) (string, error) {
	secret, err := jwtSecret()
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		user.Email,
		user.ID,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * 15)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secret)

	if err != nil {
		log.Println("Error in token signing.", err)
		return "", err
	}

	return t, nil

}

// Validate Token
func ValidateToken(clientToken string) (claims *CustomClaims, msg string) {
	secret, err := jwtSecret()
	if err != nil {
		msg = "token validation unavailable"
		return
	}

	token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		log.Printf("Token validation failed: %v", err)
		msg = "invalid token"
		return
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		msg = "invalid token claims"
		return
	}

	if !token.Valid {
		msg = "invalid token"
		return
	}

	return claims, msg
}
