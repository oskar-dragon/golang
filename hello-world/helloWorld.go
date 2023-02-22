package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefox = "Bonjour, "


func Hello(name string, language string) string {
	if (name == "") {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func main() {
	fmt.Println(Hello("Oskar", ""))
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefox
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}