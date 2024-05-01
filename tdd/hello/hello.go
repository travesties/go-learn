package hello

import (
	"fmt"
	"strings"
)

var prefixMap = map[string]string{
	"":        "Hello, ",
	"english": "Hello, ",
	"spanish": "Hola, ",
	"french":  "Bonjour, ",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	language = strings.ToLower(language)
	prefix, ok := prefixMap[language]
	if !ok {
		prefix = prefixMap[""]
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Travis", ""))
}
