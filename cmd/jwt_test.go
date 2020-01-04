package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	validJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = rescueStdout
	}()

	err := runDecodeJWT([]string{validJWT})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)

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
}
`
	if string(out) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(out))
	}
}

func TestDecodeInvalidJWT(t *testing.T) {
	validJWT := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`
	err := runDecodeJWT([]string{validJWT})
	if err.Error() != errInvalidToken {
		t.Errorf("Expected error %s, got: %s", errInvalidToken, err)
	}
}
