package cmd

import "fmt"

func green(s string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", s)
}

func red(s string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", s)
}
