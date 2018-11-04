package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"jay_medtronic/database"
	"jay_medtronic/env"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/raja/argon2pw"
)

type User struct {
	Id       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var password string

func HashPassword(plainPW string) string {
	hashedPassword, err := argon2pw.GenerateSaltedHash(plainPW)
	if err != nil {
		log.Panicf("Hash generated returned error: %v", err)
	}
	return hashedPassword
}

func ValidatePassword(hashed, plain string) (bool, error) {
	valid, err := argon2pw.CompareHashWithPassword(hashed, plain)
	return valid, err
}

func MissingFields(user User) error {
	var err error
	if user.Email == "" {
		err = errors.New("Missing email field")
	}
	if user.Password == "" {
		err = errors.New("Missing Password field")
	}
	return err
}

func CreateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email})
	// SignedString requires byte array not string as argument type
	tokenString, err := token.SignedString([]byte(env.JWTSecret))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	return tokenString
}

func ValidateToken(tokenString string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWTSecret), nil
	})
	if err != nil {
		fmt.Println("Error validating token:", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); token.Valid && ok {
		return claims["email"], true
	}
	return "Invalid token", false
}

func Create(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	hashedPW := HashPassword(user.Password)
	err = MissingFields(user)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}

	_, err = database.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, hashedPW)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	} else {
		return UserResponse{
			Message: "User created successfully",
			Token:   CreateToken(user.Email),
		}, nil
	}
}

func Login(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	err = MissingFields(user)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}

	err = database.DB.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&password)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}

	hashedPassword := password
	_, err = ValidatePassword(hashedPassword, user.Password)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}
	return UserResponse{
		Message: "User is authenticated",
		Token:   CreateToken(user.Email),
	}, nil
}
