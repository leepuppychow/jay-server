package models

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	pw := "hello"
	hashedPW := HashPassword(pw)
	valid, err := ValidatePassword(hashedPW, pw)

	if !valid || err != nil {
		t.Errorf("Validate Password failed")
	}
}
