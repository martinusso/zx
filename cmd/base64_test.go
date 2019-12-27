package cmd

import (
	"testing"
)

func TestBase64EmptyInput(t *testing.T) {
	got := processBase64([]string{})
	expected := "error: Empty input..."
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestBase64Encode(t *testing.T) {
	inputEncode = true
	got := processBase64([]string{"zx"})
	expected := "eng="
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	inputEncode = false
}

func TestBase64Decode(t *testing.T) {
	inputDecode = true
	got := processBase64([]string{"eng="})
	expected := "zx"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	inputDecode = false
}

func TestBase64DecodeError(t *testing.T) {
	inputDecode = true
	got := processBase64([]string{"zx"})
	expected := "decode error: illegal base64 data at input byte 0"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	inputDecode = false
}
