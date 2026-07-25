package main

import (
	"fmt"
)

const (
	spanishLanguage = "Spanish"
	frenchLanguage  = "French"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Holla"
	frenchHelloPrefix  = "Bonjour"
)

func prefixSelection(language string) (prefix string) {

	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := prefixSelection(language)

	return prefix + ", " + name
}

func main() {
	fmt.Println(Hello("Bruno", "English"))
}
