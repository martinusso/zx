package cmd

import (
	"fmt"
	"testing"

	"github.com/martinusso/zx/internal/cli"
)

func TestCPF(t *testing.T) {
	got := cpf([]string{})
	if len(got) != 11 {
		t.Errorf("Expected '%d', got '%d'", 11, len(got))
	}

	new := cpf([]string{got})
	expected := fmt.Sprintf("%s ➜ %s", got, cli.Green("valid"))
	if new != expected {
		t.Errorf("Expected '%s', got '%s'", expected, new)
	}
}
