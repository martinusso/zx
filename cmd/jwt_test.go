package cmd

import (
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	validJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

	got, err := runDecodeJWT([]string{validJWT})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	expected := `Header:
{
  "alg": "HS256",
  "typ": "JWT"
}

Payload:
{
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}`
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestDecodeInvalidHeaderJWT(t *testing.T) {
	invalidJWT := `eyJhbGciOiJI*zI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
	_, err := runDecodeJWT([]string{invalidJWT})
	expected := "Invalid header: illegal base64 data at input byte 12"
	if err.Error() != expected {
		t.Errorf("Expected error %s, got: %s", expected, err)
	}
}

func TestDecodeInvalidPayloadJWT(t *testing.T) {
	invalidJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIi*iIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
	_, err := runDecodeJWT([]string{invalidJWT})
	expected := "Invalid payload: illegal base64 data at input byte 8"
	if err.Error() != expected {
		t.Errorf("Expected error %s, got: %s", expected, err)
	}
}

func TestDecodeInvalidJWT(t *testing.T) {
	invalidJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`
	_, err := runDecodeJWT([]string{invalidJWT})
	if err.Error() != invalidToken {
		t.Errorf("Expected error %s, got: %s", invalidToken, err)
	}
}

func TestDecodeEmptyJWT(t *testing.T) {
	_, err := runDecodeJWT([]string{})
	if err.Error() != invalidToken {
		t.Errorf("Expected error %s, got: %s", invalidToken, err)
	}
}
