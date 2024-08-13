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
	if lang == "es" {
		return esHelloPrefix + name
	} else if lang == "fr" {
		return frHelloPrefix + name
	}
	return enHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", "en"))
}
