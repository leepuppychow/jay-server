package auth

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
