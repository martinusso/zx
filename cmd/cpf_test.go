package cmd

import (
	"fmt"
	"testing"

	"github.com/martinusso/zx/internal/cli"
)

func TestRunCPF(t *testing.T) {
	got, err := runCPF([]string{})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(got) != 11 {
		t.Errorf("Expected '%d', got '%d'", 11, len(got))
	}

	n, _ := runCPF([]string{got})
	expected := fmt.Sprintf("%s ➜ %s", got, cli.Green("valid"))
	if n != expected {
		t.Errorf("Expected '%s', got '%s'", expected, n)
	}
}
