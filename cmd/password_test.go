package cmd

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	p := generatePassword([]string{})
	if len(p) != 3 {
		t.Errorf("Expected '%d', got '%d'", 3, len(p))
	}
}

func TestGeneratePasswordWithLength(t *testing.T) {
	p := generatePassword([]string{"7"})
	if len(p) != 7 {
		t.Errorf("Expected '%d', got '%d'", 7, len(p))
	}
}

func TestGeneratePasswordWithFlag(t *testing.T) {
	passwordLength = 14
	p := generatePassword([]string{})
	if len(p) != 14 {
		t.Errorf("Expected '%d', got '%d'", 14, len(p))
	}

	passwordLength = 0
}
