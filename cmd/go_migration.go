package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cobra"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	goMigCmd = &cobra.Command{
		Use:   "gomig",
		Short: "Create migration",
		Long:  `Create migration. Args: 1 - migration path, 2 - migration name`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runGoMigUP(args)
			if err != nil {
				return err
			}
			fmt.Println(s)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(goMigCmd)
}

func runGoMigUP(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("migration args required")
	}

	path := args[0]
	name := args[1]

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	title, _, _ := transform.String(t, name)
	title = strings.Replace(title, " ", "_", -1)

	// number of nanoseconds since January 1, 1970 UTC
	nsec := time.Now().UnixNano() / int64(time.Millisecond)

	filename := fmt.Sprintf("%s/%d_%s.up.sql", path, int(nsec), title)

	err := ioutil.WriteFile(filename, []byte(""), 0777)
	if err != nil {
		return "", fmt.Errorf("Unable to write file %s: %v", filename, err)
	}
	return fmt.Sprintln("Created file:", filename), nil
}
