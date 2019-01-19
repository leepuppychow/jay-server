package models

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	email := "test@test.com"
	token := CreateToken(email)
	fmt.Println(token)
	if token == "" {
		t.Errorf("CreateToken test failed")
	}
}

func TestValidToken(t *testing.T) {
	email := "test@test.com"
	token := CreateToken(email)

	valid := ValidToken(token)
	if !valid {
		t.Errorf("ValidToken test failed")
	}
}

func TestValidatePassword(t *testing.T) {
	pw := "hello"
	hashedPW := HashPassword(pw)
	valid, err := ValidatePassword(hashedPW, pw)

	if !valid || err != nil {
		t.Errorf("Validate Password failed")
	}
}
