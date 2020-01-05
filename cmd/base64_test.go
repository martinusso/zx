package cmd

import (
	"testing"
)

func TestBase64EmptyInput(t *testing.T) {
	got, err := runBase64([]string{})
	if err.Error() != errEmptyInput {
		t.Errorf("Expected '%s', got '%s'", errEmptyInput, err)
	}
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestBase64Encode(t *testing.T) {
	inputEncode = true
	got, err := runBase64([]string{"zx"})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := "eng="
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	inputEncode = false
}

func TestBase64Decode(t *testing.T) {
	inputDecode = true
	got, err := runBase64([]string{"eng="})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := "zx"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	inputDecode = false
}

func TestBase64DecodeError(t *testing.T) {
	inputDecode = true
	_, err := runBase64([]string{"zx"})
	expectedErr := "Decode error: illegal base64 data at input byte 0"
	if err.Error() != expectedErr {
		t.Errorf("Expected '%s', got '%s'", expectedErr, err)
	}

	inputDecode = false
}
