package main

import "fmt"

func Hello(s string) string {
	return "hello," + s
}

func main() {
	fmt.Println(Hello("chris"))
}
