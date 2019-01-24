package models

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ValidToken(tokenString string) bool {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error validating token or invalid token:", err)
		return false
	}
	return true
}

func CreateToken(email string) string {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	return tokenString
}
