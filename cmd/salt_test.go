package cmd

import (
	"strings"
	"testing"
)

const (
	hashLength   = 128
	passwordTest = "Qwerty!@#456"
)

func TestGenerateSalt(t *testing.T) {
	s, err := runSalt([]string{passwordTest})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if !strings.Contains(s, passwordTest) {
		t.Errorf("Expected contains password '%s' in '%s'", passwordTest, s)
	}
	expected := 24 + 12 + hashLength + hashLength
	if len(s) != expected {
		t.Errorf("Expected '%d', got '%d'", expected, len(s))
	}
}

func TestGenerateSaltWithoutPassword(t *testing.T) {
	s, err := runSalt([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := 24 + 8 + hashLength + hashLength
	if len(s) != expected {
		t.Errorf("Expected '%d', got '%d'", expected, len(s))
	}
}

func TestGenerateSaltWithPasswordFlag(t *testing.T) {
	saltPassword = passwordTest
	s, err := runSalt([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if !strings.Contains(s, passwordTest) {
		t.Errorf("Expected contains password '%s' in '%s'", passwordTest, s)
	}
	expected := 24 + 12 + hashLength + hashLength
	if len(s) != expected {
		t.Errorf("Expected '%d', got '%d'", expected, len(s))
	}

	saltPassword = ""
}
