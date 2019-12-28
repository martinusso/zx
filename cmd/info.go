package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	zxDescription = `zx is a set of handy commands to make some daily tasks easier.`
	zxVersion     = "v0.2.0"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the info of zx",
	Long:  `Print the info of zx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(zxDescription + `
version: ` + zxVersion)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
