package main

import (
	"fmt"
	"jay_medtronic/models"
	"testing"
)

func TestCreateToken(t *testing.T) {
	email := "test@test.com"
	token := user.CreateToken(email)
	fmt.Println(token)
	if token == "" {
		t.Errorf("CreateToken test failed")
	}
}

func TestValidateToken(t *testing.T) {
	email := "test@test.com"
	token := user.CreateToken(email)

	emailFromToken, valid := user.ValidateToken(token)
	if !valid || emailFromToken != "test@test.com" {
		t.Errorf("ValidateToken test failed")
	}
}

func TestValidatePassword(t *testing.T) {
	pw := "hello"
	hashedPW := user.HashPassword(pw)
	valid, err := user.ValidatePassword(hashedPW, pw)

	if !valid || err != nil {
		t.Errorf("Validate Password failed")
	}
}
