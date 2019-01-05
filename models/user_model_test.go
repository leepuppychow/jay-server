package models_test

import (
	"fmt"
	"testing"

	"github.com/leepuppychow/jay_medtronic/models"
)

func TestCreateToken(t *testing.T) {
	email := "test@test.com"
	token := models.CreateToken(email)
	fmt.Println(token)
	if token == "" {
		t.Errorf("CreateToken test failed")
	}
}

func TestValidToken(t *testing.T) {
	email := "test@test.com"
	token := models.CreateToken(email)

	valid := models.ValidToken(token)
	if !valid {
		t.Errorf("ValidToken test failed")
	}
}

func TestValidatePassword(t *testing.T) {
	pw := "hello"
	hashedPW := models.HashPassword(pw)
	valid, err := models.ValidatePassword(hashedPW, pw)

	if !valid || err != nil {
		t.Errorf("Validate Password failed")
	}
}
