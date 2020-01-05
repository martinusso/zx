package cmd

import (
	"fmt"

	helper "github.com/martinusso/go-docs/cpf"
	"github.com/martinusso/zx/internal/cli"
	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

var (
	cpfCmd = &cobra.Command{
		Use:   "cpf",
		Short: "Generate/Validate CPF",
		Long:  `Generate a valid CPF or Validate if pass a CPF as args`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runCPF(args)
			if err != nil {
				return err
			}
			clipboard.Write(s)
			fmt.Println(s)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(cpfCmd)
}

func runCPF(args []string) (string, error) {
	var cpf string
	if len(args) > 0 {
		cpf = args[0]
	}
	if cpf != "" {
		return validateCPF(cpf), nil
	}
	return generateCPF(), nil
}

func generateCPF() string {
	return helper.Generate()
}

func validateCPF(cpf string) string {
	v := cli.Red("invalid")
	if helper.Valid(cpf) {
		v = cli.Green("valid")
	}
	return fmt.Sprintf("%s ➜ %s", cpf, v)
}
