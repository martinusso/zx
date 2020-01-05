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

func TestDecodeInvalidJWT(t *testing.T) {
	validJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`
	_, err := runDecodeJWT([]string{validJWT})
	if err.Error() != invalidToken {
		t.Errorf("Expected error %s, got: %s", invalidToken, err)
	}
}
