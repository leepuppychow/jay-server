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
	Username string `json:"username"`
	Password string `json:"password"`
}

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

func Create(body io.Reader) (string, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)

	if user.Username == "" {
		err = errors.New("Missing Username field")
	}
	if user.Password == "" {
		err = errors.New("Missing Password field")
	}

	hashedPW := HashPassword(user.Password)

	if err != nil {
		return "Error creating user", err
	} else {
		_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPW)
		return "User created successfully", nil
	}
}
