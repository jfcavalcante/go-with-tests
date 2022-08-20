package main

import "fmt"

// func main() {
// Side effect = printing to stdout
// fmt.Println("Hello, world") // Current domain -> the string we send t outside
// }

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "
const portugueseOlaPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// named return value
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHolaPrefix
	case french:
		prefix = frenchBonjourPrefix
	case portuguese:
		prefix = portugueseOlaPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
