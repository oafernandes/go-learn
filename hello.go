package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "world!"
	}

	return greetingPrefix(language) + name
}

//public functions start upper, private start lower
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	//default is used if none of the other case statements match
	default:
		prefix = englishHelloPrefix
	}
	//returns prefix which was created in the function signature
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
