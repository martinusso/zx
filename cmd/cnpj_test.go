package cmd

import (
	"fmt"
	"testing"

	"github.com/martinusso/zx/internal/cli"
)

func TestCNPJ(t *testing.T) {
	got, err := runCNPJ([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(got) != 14 {
		t.Errorf("Expected '%d', got '%d'", 14, len(got))
	}

	n, err := runCNPJ([]string{got})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := fmt.Sprintf("%s ➜ %s", got, cli.Green("valid"))
	if n != expected {
		t.Errorf("Expected '%s', got '%s'", expected, n)
	}
}
