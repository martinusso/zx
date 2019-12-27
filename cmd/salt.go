package cmd

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	saltPassword string

	saltCmd = &cobra.Command{
		Use:   "salt",
		Short: "Generate a random and unique salt and hash for passwords",
		Long:  `Generate a random and unique salt and hash for passwords`,
		Run: func(cmd *cobra.Command, args []string) {
			p, h, s := generateSalt(args)
			fmt.Println(fmt.Sprintf("Password: %s\nHash: %s\nSalt: %s", p, h, s))
		},
	}
)

func init() {
	saltCmd.Flags().StringVarP(&saltPassword, "password", "p", "", "Password")
	rootCmd.AddCommand(saltCmd)
}

func generateSalt(args []string) (p string, h string, s string) {
	p = getPassword(args)
	s = getSalt()
	h = hash(p + s)
	return
}

func getPassword(args []string) string {
	if saltPassword != "" {
		return saltPassword
	}
	if len(args) > 0 {
		return args[0]
	}

	return generatePassword([]string{"8"})
}

func getSalt() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	r := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return hash(r)
}

func hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
