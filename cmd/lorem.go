package cmd

import (
	"fmt"
	"strconv"

	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

const (
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris in lorem ac elit blandit suscipit nec eu mi. In et nisi a odio tempus ultrices. Duis dapibus nisi odio, ut tristique justo pellentesque at. Sed euismod erat eros, eget varius erat gravida id. Quisque euismod fermentum placerat. Curabitur ac varius quam. Ut a lorem at urna accumsan consequat ac sit amet enim. Nullam nunc lorem, efficitur ut scelerisque vel, porta quis tortor. Cras a sagittis ante. Aliquam pharetra ex sem, sit amet vulputate est commodo ut.`
)

var (
	loremCmd = &cobra.Command{
		Use:   "lorem",
		Short: "Generate Lorem Ipsum",
		Long:  `Generate Lorem Ipsum`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := generateLorem(args)
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
	rootCmd.AddCommand(loremCmd)
}

func generateLorem(args []string) (string, error) {
	var size int
	if len(args) > 0 {
		size, _ = strconv.Atoi(args[0])
	}
	if size <= 0 || size > len(lorem) {
		return lorem, nil
	}
	return lorem[0:size], nil
}
