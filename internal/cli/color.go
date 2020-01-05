package cli

import "fmt"

func Cyan(s string) string {
	return fmt.Sprintf("\036[1;32m%s\036[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", s)
}

func Magenta(s string) string {
	return fmt.Sprintf("\035[1;31m%s\035[0m", s)
}

func Red(s string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", s)
}
