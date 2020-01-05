package cmd

import (
	"fmt"

	"github.com/martinusso/zx/internal/web"
	"github.com/spf13/cobra"
)

type erResponse struct {
	Real struct {
		BRL float32 `json:"BRL"`
		USD float32 `json:"USD"`
		EUR float32 `json:"EUR"`
		GBP float32 `json:"GBP"`
	} `json:"rates"`
}

var (
	url = "https://api.exchangeratesapi.io/latest?base=USD"

	exchangeRateCmd = &cobra.Command{
		Use:   "exchange",
		Short: "List of foreign currency rates",
		Long: `List of foreign currency rates.
Reference exchange rate: US dollar (USD).
BRL: ? (BRL 1 = USD ?), EUR: ? (EUR 1 = USD ?), GBP: ? (GBP 1 = USD ?)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runExchangeRate()
			if err != nil {
				return err
			}
			fmt.Println(s)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(exchangeRateCmd)
}

func runExchangeRate() (string, error) {
	data := &erResponse{}
	err := web.GetJSON(url, data)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("BRL: %.2f (USD %.2f), EUR: %.2f (USD %.2f), GBP: %.2f (USD %.2f)",
		data.Real.BRL,
		1/data.Real.BRL,
		data.Real.EUR,
		1/data.Real.EUR,
		data.Real.GBP,
		1/data.Real.GBP)
	return s, nil
}
