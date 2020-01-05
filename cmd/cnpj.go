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
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runCNPJ(args)
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
	rootCmd.AddCommand(cnpjCmd)
}

func runCNPJ(args []string) (string, error) {
	var cnpj string
	if len(args) > 0 {
		cnpj = args[0]
	}
	if cnpj != "" {
		return validateCNPJ(cnpj), nil
	}
	return generateCNPJ(), nil
}

func generateCNPJ() string {
	return helper.Generate()
}

func validateCNPJ(cnpj string) string {
	v := cli.Red("invalid")
	if helper.Valid(cnpj) {
		v = cli.Green("valid")
	}
	return fmt.Sprintf("%s ➜ %s", cnpj, v)
}
