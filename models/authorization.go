package models

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/leepuppychow/jay_medtronic/env"
)

func ValidToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error validating token or invalid token:", err)
		return false
	}
	return true
}

func CreateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email})
	tokenString, err := token.SignedString([]byte(env.JWTSecret))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	return tokenString
}
