package main

import (
	"fmt"
)

const (
	enHelloPrefix = "Hello "
	esHelloPrefix = "Hola "
	frHelloPrefix = "Bonjour "
)

func Hello(name, lang string) string {
	if name == "" {
		if lang == "es" {
			name = "mundo"
		} else if lang == "fr" {
			name = "monde"
		} else {
			name = "world"
		}
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "es":
		prefix = esHelloPrefix
	case "fr":
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", "en"))
}
