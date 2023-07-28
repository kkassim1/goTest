package main

import "fmt"

const (
	englishHelloePreflix = "hello,"
	spanish              = "spanish"
	spanishprefix        = "hola,"
	helloworld           = "Hello World"
	frenchpreflix        = "bonjour,"
	french               = "french"
	english              = "english"
)

func Hello(name, language string) string {
	if name == "" {
		return helloworld
	}

	if language == english {
		return englishHelloePreflix + name
	}

	if language == spanish {
		return spanishprefix + name
	}
	if language == french {
		return frenchpreflix + name
	}
	return englishHelloePreflix + name
}

func main() {
	fmt.Println(Hello("chris", ""))
}
