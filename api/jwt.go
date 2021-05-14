package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createJWTtoken(id interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().AddDate(0, 1, 0).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", fmt.Errorf("cannot get signed token string: %v", err)
	}

	return tokenString, nil
}
