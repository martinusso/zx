package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	language string
)

var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "Mock generates mock interfaces from a source file",
	Long:  `Mock generates mock interfaces from a source file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := runMock(args)
		if err != nil {
			return err
		}
		fmt.Println(s)
		return nil
	},
}

func init() {
	mockCmd.Flags().StringVarP(&language, "language", "l", "go", "Programming language")
	rootCmd.AddCommand(mockCmd)
}

func runMock(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(emptyInput)
	}
	filepath := completePath(args[0])
	c, a := goCmd(filepath)
	fmt.Printf("Running `%s %s`\n", c, strings.Join(a, " "))
	err := exec.Command(c, a...).Run()
	if err != nil {
		return "", err
	}
	return "Done!", nil
}

func goCmd(name string) (string, []string) {
	dest := strings.Replace(name, ".go", "_mock.go", 1)
	pkg := filepath.Base(filepath.Dir(name))
	args := []string{
		fmt.Sprintf("-source=%s", name),
		fmt.Sprintf("-destination=%s", dest),
		fmt.Sprintf("-package=%s", pkg),
	}
	return "mockgen", args
}

func completePath(filepath string) string {
	if filepath[0] == '/' {
		return filepath
	}
	dir, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", dir, filepath)
}
