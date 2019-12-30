package cmd

import (
	"fmt"

	"github.com/martinusso/zx/internal/web"
	"github.com/spf13/cobra"
)

type erResponse struct {
	Real struct {
		USD float32 `json:"USD"`
		EUR float32 `json:"EUR"`
		GBP float32 `json:"GBP"`
	} `json:"rates"`
}

var (
	url = "https://api.exchangeratesapi.io/latest?base=BRL"

	exchangeRateCmd = &cobra.Command{
		Use:   "exchange",
		Short: "List of foreign currency rates",
		Long:  `List of foreign currency rates`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(exchangeRate())
		},
	}
)

func init() {
	rootCmd.AddCommand(exchangeRateCmd)
}

func exchangeRate() string {
	data := &erResponse{}
	err := web.GetJSON(url, data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Dólar: %.2f, Euro: %.2f, Libra: %.2f",
		1/data.Real.USD,
		1/data.Real.EUR,
		1/data.Real.GBP)
}
