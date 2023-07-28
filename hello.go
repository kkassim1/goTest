package main

import "fmt"

const englishHelloePreflix = "Hello,"

func Hello(name string) string {
	if name == "" {
		return "Hello World"
	}

	return englishHelloePreflix + name
}

func main() {
	fmt.Println(Hello("chris"))
}
