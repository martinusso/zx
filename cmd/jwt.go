package cmd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	invalidToken = "Invalid token: token should contain header, payload and secret"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Decode a JWT token",
	Long: `Decode a JWT token.
This command doesn't validate the token, any well formed JWT can be decoded.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := runDecodeJWT(args)
		if err != nil {
			return err
		}
		fmt.Println(s)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

func runDecodeJWT(args []string) (string, error) {
	token, err := getTokenJWT(args)
	if err != nil {
		return "", err
	}

	header, err := decodeSegment(token[0])
	if err != nil {
		return "", fmt.Errorf("Invalid header: %s", err.Error())
	}

	payload, err := decodeSegment(token[1])
	if err != nil {
		return "", fmt.Errorf("Invalid payload: %s", err.Error())
	}

	return fmt.Sprintf("Header:\n%s\n\nPayload:\n%s", header, payload), nil
}

func getTokenJWT(args []string) (token []string, err error) {
	if len(args) > 0 {
		token = strings.Split(args[0], ".")
	}

	// check if the jwt token contains header, payload and token
	if len(token) != 3 {
		err = errors.New(invalidToken)
	}
	return
}

// Decode JWT specific base64url encoding with padding stripped
func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(seg)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	json.Unmarshal(decoded, &data)
	return json.MarshalIndent(data, "", "  ")
}
