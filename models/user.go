package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/leepuppychow/jay-server/auth"
	"github.com/leepuppychow/jay-server/database"
	"github.com/raja/argon2pw"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
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

func CreateUser(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	err = MissingFields(user)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}

	hashedPW := HashPassword(user.Password)
	_, err = database.DB.Exec("INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)", user.Email, hashedPW)

	if err != nil {
		fmt.Println(err)
		return UserResponse{Message: "Unable to create user"}, err
	} else {
		return UserResponse{
			Message: "User created successfully",
			Token:   auth.CreateToken(user.Email),
		}, nil
	}
}

func LoginUser(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	err = MissingFields(user)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}

	err = database.DB.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&password)

	if err != nil {
		return UserResponse{Message: "Unable to find user"}, err
	}

	hashedPassword := password
	_, err = ValidatePassword(hashedPassword, user.Password)

	if err != nil {
		return UserResponse{Message: err.Error()}, err
	}
	return UserResponse{
		Message: "User is authenticated",
		Token:   auth.CreateToken(user.Email),
	}, nil
}
