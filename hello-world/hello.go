package hello

import (
	"fmt"
)

func main() {
	fmt.Println("Olá Mundo")
}

const portuguese_prefix = "Olá"
const spanish_prefix = "Hola"
const english_prefix = "Hello"

func Ola(name string, language string) string {
	if name == "" {
		return cratePrefix(language)
	}
	return cratePrefix(language) + " " + name
}

func cratePrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanish_prefix
	case "portuguese":
		prefix = portuguese_prefix
	default:
		prefix = english_prefix
	}
	return
}
