package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestCompletePath(t *testing.T) {
	dir, _ := os.Getwd()
	m := map[string]string{
		"mockgen.go":  fmt.Sprintf("%s/mockgen.go", dir),
		"/mockgen.go": "/mockgen.go",
	}
	for k, v := range m {
		got := completePath(k)
		if got != v {
			t.Errorf("Expected '%s', got '%s'", v, got)
		}
	}
}
