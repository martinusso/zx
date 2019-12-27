package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	zxDescription = `zx is a set of handy commands to make some daily tasks easier and more fun.`
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the info of zx",
	Long:  `Print the info of zx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(zxDescription + `
version: v0.1.0`)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
