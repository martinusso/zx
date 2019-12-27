package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inputEncode bool
	inputDecode bool

	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Decode from Base64 or Encode to Base64",
		Long: `Decode from Base64 or Encode to Base64
zx base64 -e zx
zx base64 -d eng=`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(processBase64(args))
		},
	}
)

func init() {
	base64Cmd.Flags().BoolVarP(&inputDecode, "decode", "d", false, "Decode from Base64")
	base64Cmd.Flags().BoolVarP(&inputEncode, "encode", "e", false, "Encode to Base64")
	rootCmd.AddCommand(base64Cmd)
}

func processBase64(args []string) string {
	if len(args) == 0 {
		return "error: Empty input..."
	}

	s := args[0]

	// encode
	if isEncode() {
		return base64.StdEncoding.EncodeToString([]byte(s))
	}

	// decode
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Sprintf("decode error: %s", err)
	}
	return string(decoded)
}

func isEncode() bool {
	return inputEncode
}
