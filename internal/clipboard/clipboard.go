package clipboard

import cb "github.com/atotto/clipboard"

func Write(s string) {
	cb.WriteAll(s)
}
