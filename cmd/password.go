package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const (
	letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits   = "0123456789"
	specials = "~=+%^*/()[]{}/!@#$?|"
)

var (
	passwordLength int

	passwordCmd = &cobra.Command{
		Use:   "password [SIZE]",
		Short: "Generate a random password",
		Long:  `Generate a random password`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(generatePassword(args))
		},
	}
)

func init() {
	passwordCmd.Flags().IntVarP(&passwordLength, "length", "l", 0, "Password length")
	rootCmd.AddCommand(passwordCmd)
}

func generatePassword(args []string) string {
	rand.Seed(time.Now().UnixNano())

	length := getPasswordLength(args)
	all := letters + digits + specials

	buf := make([]byte, length)
	buf[0] = letters[rand.Intn(len(letters))]
	buf[1] = digits[rand.Intn(len(digits))]
	buf[2] = specials[rand.Intn(len(specials))]
	for i := 3; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func getPasswordLength(args []string) (length int) {
	if passwordLength != 0 {
		length = passwordLength
	} else if len(args) > 0 {
		length, _ = strconv.Atoi(args[0])
	}
	if length < 3 {
		length = 3
	}
	return
}
