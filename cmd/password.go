package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/martinusso/zx/internal/clipboard"
	"github.com/spf13/cobra"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	symbols      = "~=+%^*/()[]{}/!@#$?|"
)

var (
	inputLength int
	inputNo     string

	passwordCmd = &cobra.Command{
		Use:   "password",
		Short: "Generate a random password",
		Long:  `Generate a random password`,
		Run: func(cmd *cobra.Command, args []string) {
			s := generatePassword(args)
			clipboard.Write(s)
			fmt.Println(s)
		},
	}
)

func init() {
	passwordCmd.Flags().IntVarP(&inputLength, "length", "l", 0, "Password length")
	passwordCmd.Flags().StringVarP(&inputNo, "no", "N", "", "Password without [l]owercase, [u]ppercase, [n]umbers, [s]ymbols")
	rootCmd.AddCommand(passwordCmd)
}

func generatePassword(args []string) string {
	length := getPasswordLength(args)
	return makePassword(length)
}

func makePassword(l int) string {
	var all string
	var i int

	buf := make([]byte, l)
	if !strings.Contains(inputNo, "l") {
		buf[i] = lowerLetters[rand.Intn(len(lowerLetters))]
		all += lowerLetters
		i += 1
	}
	if !strings.Contains(inputNo, "u") {
		buf[i] = upperLetters[rand.Intn(len(upperLetters))]
		all += upperLetters
		i += 1
	}
	if !strings.Contains(inputNo, "n") {
		buf[i] = numbers[rand.Intn(len(numbers))]
		all += numbers
		i += 1
	}
	if !strings.Contains(inputNo, "s") {
		buf[i] = symbols[rand.Intn(len(symbols))]
		all += symbols
		i += 1
	}

	if all == "" {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	for i := i; i < l; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func getPasswordLength(args []string) (length int) {
	if inputLength != 0 {
		length = inputLength
	} else if len(args) > 0 {
		length, _ = strconv.Atoi(args[0])
	}
	if length < 3 {
		length = 8
	}
	return
}
