package cmd

import (
	"fmt"
	"testing"

	"github.com/martinusso/zx/internal/cli"
)

func TestCNPJ(t *testing.T) {
	got := cnpj([]string{})
	if len(got) != 14 {
		t.Errorf("Expected '%d', got '%d'", 14, len(got))
	}

	new := cnpj([]string{got})
	expected := fmt.Sprintf("%s ➜ %s", got, cli.Green("valid"))
	if new != expected {
		t.Errorf("Expected '%s', got '%s'", expected, new)
	}
}
