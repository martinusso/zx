package clipboard

import cb "github.com/atotto/clipboard"

func Write(s string) {
	if s == "" {
		return
	}
	cb.WriteAll(s)
}
