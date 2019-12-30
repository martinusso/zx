package cmd

import (
	"fmt"

	uuid "github.com/beevik/guid"
	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

var (
	uuidPassword string

	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generate a random UUID",
		Long:  `Generate a random UUID`,
		Run: func(cmd *cobra.Command, args []string) {
			uuid := generateUUID()
			clipboard.Write(uuid)
			fmt.Println(uuid)
		},
	}
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

func generateUUID() string {
	return uuid.NewString()
}
