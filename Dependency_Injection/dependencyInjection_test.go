package dependencyinjection

import (
	"testing"
	"bytes"
	"os"
)

func TestGreet(t * testing.T) {

	t.Run("Test Greet 1 : Use buffer as input writter", func(t * testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	})

	t.Run("Test Greet 2 : Use Stdout as output writter", func(t * testing.T) {
		bytesNumber, _ := Greet(os.Stdout, "Chris")
		want := len("Hello, Chris")

		if bytesNumber != want {
			t.Errorf("bytesNumber %v != want %d", bytesNumber, want)
		}
	})
}