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
	ipURL         = "https://api.db-ip.com/v2/free"
	urlWhatIsMyIP = "https://api.ipify.org?format=json"

	ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "IP geolocation and what is my IP address",
		Long: `IP geolocation. [How to use: zx ip 187.36.173.130]
It provides a simple IP to country, state and city mapping. Using db-ip.com/api.

What is my IP address. [How to use: zx ip my]
It shows your IPv4 address. Using ipify.org API.`,
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
	if inputIP == "my" {
		return whatIsMyIP()
	}
	return ipGeolocation(inputIP)
}

func ipGeolocation(inputIP string) (string, error) {
	data := &ipResponse{}

	url := fmt.Sprintf("%s/%s", ipURL, inputIP)
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

func whatIsMyIP() (string, error) {
	var data struct {
		IP string `json:"ip"`
	}

	err := web.GetJSON(urlWhatIsMyIP, &data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Your IPv4 Address Is: %s\n", data.IP), nil
}
