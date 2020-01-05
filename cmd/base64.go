package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

var (
	inputEncode bool
	inputDecode bool
)

var base64Cmd = &cobra.Command{
	Use:     "base64",
	Aliases: []string{"b64"},
	Short:   "Decode from Base64 or Encode to Base64",
	Long:    `Decode from Base64 or Encode to Base64.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := runBase64(args)
		if err != nil {
			return err
		}
		clipboard.Write(s)
		fmt.Println(s)
		return nil
	},
}

func init() {
	base64Cmd.Flags().BoolVarP(&inputDecode, "decode", "d", false, "Decode from Base64")
	base64Cmd.Flags().BoolVarP(&inputEncode, "encode", "e", false, "Encode to Base64")
	rootCmd.AddCommand(base64Cmd)
}

func runBase64(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New(emptyInput)
	}

	input := args[0]

	// encode
	if inputEncode {
		return base64.StdEncoding.EncodeToString([]byte(input)), nil
	}

	// decode
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("Decode error: %s", err.Error())
	}
	return string(decoded), nil
}
