package cmd

import (
	"errors"
	"fmt"

	"github.com/martinusso/zx/internal/clipboard"
	"github.com/martinusso/zx/internal/web"
	"github.com/spf13/cobra"
)

type ipResponse struct {
	IP            string `json:"ipAddress"`
	ContinentCode string `json:"continentCode"`
	ContinentName string `json:"continentName"`
	CountryCode   string `json:"countryCode"`
	CountryName   string `json:"countryName"`
	StateProv     string `json:"stateProv"`
	City          string `json:"city"`
}

var (
	ipURL = "https://api.db-ip.com/v2/free"

	ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "IP geolocation",
		Long: `IP geolocation.
It provides a simple IP to country, state and city mapping. Using db-ip.com/api.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runIP(args)
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
	rootCmd.AddCommand(ipCmd)
}

func runIP(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(emptyInput)
	}
	inputIP := args[0]

	url := fmt.Sprintf("%s/%s", ipURL, inputIP)

	data := &ipResponse{}
	err := web.GetJSON(url, data)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("IP: %s\nContinent: %s (%s)\nCountry: %s (%s)\nCity: %s (%s)",
		data.IP,
		data.ContinentName,
		data.ContinentCode,
		data.CountryName,
		data.CountryCode,
		data.City,
		data.StateProv)
	return s, nil
}
