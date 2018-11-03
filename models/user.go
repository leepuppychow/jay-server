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
	Email string `json:"email"`
	Password string `json:"password"`
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

func Create(body io.Reader) (string, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		log.Fatal("Error decoding body into JSON")
	}
	err = MissingFields(user)
	hashedPW := HashPassword(user.Password)

	if err != nil {
		return "Error creating user", err
	} else {
		_, err = database.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, hashedPW)
		return "User created successfully", nil
	}
}

func Login(body io.Reader) (string, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		log.Fatal("Error decoding body into JSON")
	}
	err = MissingFields(user)
	if err != nil {
		return "Missing fields", err
	}
	err = database.DB.QueryRow("SELECT password FROM users WHERE email=$1", user.Email).Scan(&password)
	if err != nil {
		return "Error searching for user", err
	}
	hashedPassword := password
	_, err = ValidatePassword(hashedPassword, user.Password)
	if err != nil {
		return "Error with argon2 password validation", err
	}
	return "User is authenticated", nil
}
