package cmd

import (
	"fmt"

	helper "github.com/martinusso/go-docs/cpf"
	"github.com/spf13/cobra"
)

var (
	cpfCmd = &cobra.Command{
		Use:   "cpf",
		Short: "Generate/Validate CPF",
		Long:  `Generate a valid CPF or Validate if pass a CPF as args`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cpf(args))
		},
	}
)

func init() {
	rootCmd.AddCommand(cpfCmd)
}

func cpf(args []string) string {
	var cpf string
	if len(args) > 0 {
		cpf = args[0]
	}
	if cpf != "" {
		return validCPF(cpf)
	}
	return generateCPF()
}

func generateCPF() string {
	return helper.Generate()
}

func validCPF(cpf string) string {
	v := red("invalid")
	if helper.Valid(cpf) {
		v = green("valid")
	}
	return fmt.Sprintf("%s ➜ %s", cpf, v)
}
