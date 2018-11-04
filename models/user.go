package user

import (
	"encoding/json"
	"errors"
	"io"
	"jay_medtronic/database"
	"log"

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

func Create(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	hashedPW := HashPassword(user.Password)
	err = MissingFields(user)
	
	if err != nil {
		return UserResponse{ Message: err.Error() }, err
	} 

	_, err = database.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, hashedPW)
	if err != nil {
		return UserResponse{ Message: err.Error() }, err
	} else {
		return UserResponse{
			Message: "User created successfully",
			Token:   "WHEEE",
		}, nil
	}
}

func Login(body io.Reader) (UserResponse, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	err = MissingFields(user)

	if err != nil {
		return UserResponse{ Message: err.Error() }, err
	}

	err = database.DB.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&password)
	if err != nil {
		return UserResponse{ Message: err.Error() }, err
	}
	hashedPassword := password
	_, err = ValidatePassword(hashedPassword, user.Password)
	if err != nil {
		return UserResponse{ Message: err.Error() }, err
	}
	return UserResponse{
		Message: "User is authenticated",
		Token:   "WHEEE",
	}, nil
}
