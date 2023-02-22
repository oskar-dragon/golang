package helloWorld

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefox = "Bonjour, "


func Hello(name string, language string) string {
	if (name == "") {
		name = "World"
	}

	prefix := GreetingPrefix(language)

	return prefix + name
}


func GreetingPrefix(language string) (prefix string) {
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