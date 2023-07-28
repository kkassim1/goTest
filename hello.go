package main

import "fmt"

const (
	englishHelloePreflix = "hello,"
	spanishprefix        = "hola,"
	frenchpreflix        = "bonjour,"
	yorubaprefix         = "báwo,"

	spanish    = "spanish"
	helloworld = "Hello World"
	french     = "french"
	english    = "english"
	yoruba     = "yoruba"
)

func main() {
	fmt.Println(Hello("chris", ""))
}

func Hello(name, language string) string {

	if name == "" {
		return helloworld
	}

	return greetingprefix(language) + name
}

func greetingprefix(language string) (prefix string) {

	switch language {
	case english:
		prefix = englishHelloePreflix
	case spanish:
		prefix = spanishprefix
	case french:
		prefix = frenchpreflix
	case yoruba:
		prefix = yorubaprefix
	default:
		prefix = englishHelloePreflix
	}

	return
}
