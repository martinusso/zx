package cli

import "fmt"

func Green(s string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", s)
}

func Red(s string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", s)
}
