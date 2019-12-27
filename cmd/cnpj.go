package cmd

import (
	"fmt"

	helper "github.com/martinusso/go-docs/cnpj"
	"github.com/spf13/cobra"
)

var (
	cnpjCmd = &cobra.Command{
		Use:   "cnpj",
		Short: "Generate/Validate CNPJ",
		Long:  `Generate a valid CNPJ or Validate if pass a CNPJ as args`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cnpj(args))
		},
	}
)

func init() {
	rootCmd.AddCommand(cnpjCmd)
}

func cnpj(args []string) string {
	var cnpj string
	if len(args) > 0 {
		cnpj = args[0]
	}
	if cnpj != "" {
		return validCNPJ(cnpj)
	}
	return generateCNPJ()
}

func generateCNPJ() string {
	return helper.Generate()
}

func validCNPJ(cnpj string) string {
	v := red("invalid")
	if helper.Valid(cnpj) {
		v = green("valid")
	}
	return fmt.Sprintf("%s ➜ %s", cnpj, v)
}
