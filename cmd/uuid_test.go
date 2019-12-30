package cmd

import (
	"testing"
)

func TestMethodName(t *testing.T) {
	got := generateUUID()
	expectedLength := 36
	if len(got) != expectedLength {
		t.Errorf("Expected '%d', got '%d'", expectedLength, len(got))
	}
}
