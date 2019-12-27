package cmd

import (
	"testing"
)

const (
	hashLength   = 128
	passwordTest = "Qwerty!@#456"
)

func TestGenerateSalt(t *testing.T) {
	p, h, s := generateSalt([]string{passwordTest})
	if p != passwordTest {
		t.Errorf("Expected '%s', got '%s'", passwordTest, p)
	}
	if len(h) != hashLength {
		t.Errorf("Expected '%d', got '%d'", hashLength, len(h))
	}
	if len(s) != hashLength {
		t.Errorf("Expected '%d', got '%d'", hashLength, len(s))
	}
}

func TestGenerateSaltWithoutPassword(t *testing.T) {
	p, _, _ := generateSalt([]string{})
	if len(p) != 8 {
		t.Errorf("Expected '%d', got '%d'", 8, len(p))
	}
}

func TestGenerateSaltWithPasswordFlag(t *testing.T) {
	saltPassword = passwordTest
	p, _, _ := generateSalt([]string{})
	if p != passwordTest {
		t.Errorf("Expected '%s', got '%s'", passwordTest, p)
	}

	saltPassword = ""
}
