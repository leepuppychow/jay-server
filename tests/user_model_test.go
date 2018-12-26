package main

import (
	"fmt"
	"jay_medtronic/models"
	"testing"
)

func TestCreateToken(t *testing.T) {
	email := "test@test.com"
	token := models.CreateToken(email)
	fmt.Println(token)
	if token == "" {
		t.Errorf("CreateToken test failed")
	}
}

func TestValidateToken(t *testing.T) {
	email := "test@test.com"
	token := models.CreateToken(email)

	emailFromToken, valid := models.ValidateToken(token)
	if !valid || emailFromToken != "test@test.com" {
		t.Errorf("ValidateToken test failed")
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
