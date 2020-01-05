package cmd

import (
	"testing"
)

func TestRunPassword(t *testing.T) {
	p, err := runPassword([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(p) != 8 {
		t.Errorf("Expected '%d', got '%d'", 8, len(p))
	}
}

func TestRunPasswordWithLength(t *testing.T) {
	p, err := runPassword([]string{"7"})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(p) != 7 {
		t.Errorf("Expected '%d', got '%d'", 7, len(p))
	}
}

func TestRunPasswordWithFlag(t *testing.T) {
	inputLength = 14
	p, err := runPassword([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(p) != 14 {
		t.Errorf("Expected '%d', got '%d'", 14, len(p))
	}

	inputLength = 0
}
