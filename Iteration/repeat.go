package iteration

import "strings"

func Repeat(occurences int, word string) string {
	var repeated strings.Builder

	for i := 0; i < occurences; i++ {
		repeated.WriteString(word)
	}

	return repeated.String()
}
