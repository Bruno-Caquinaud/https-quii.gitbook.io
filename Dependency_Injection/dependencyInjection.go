package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(input io.Writer, word string) (int, error) {
	return fmt.Fprintf(input, "Hello, %s", word)
}