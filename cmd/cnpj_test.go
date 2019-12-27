package cmd

import (
	"fmt"
	"testing"
)

func TestCNPJ(t *testing.T) {
	got := cnpj([]string{})
	if len(got) != 14 {
		t.Errorf("Expected '%d', got '%d'", 14, len(got))
	}

	new := cnpj([]string{got})
	expected := fmt.Sprintf("%s ➜ %s", got, green("valid"))
	if new != expected {
		t.Errorf("Expected '%s', got '%s'", expected, new)
	}
}
