package cmd

import (
	"fmt"

	helper "github.com/martinusso/go-docs/cnpj"
	"github.com/martinusso/zx/internal/cli"
	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

var (
	cnpjCmd = &cobra.Command{
		Use:   "cnpj",
		Short: "Generate/Validate CNPJ",
		Long:  `Generate a valid CNPJ or Validate if pass a CNPJ as args`,
		Run: func(cmd *cobra.Command, args []string) {
			s := cnpj(args)
			clipboard.Write(s)
			fmt.Println(s)
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
	v := cli.Red("invalid")
	if helper.Valid(cnpj) {
		v = cli.Green("valid")
	}
	return fmt.Sprintf("%s ➜ %s", cnpj, v)
}
