package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetPrefix(language) + name
}

func greetPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Atanda", ""))
}
