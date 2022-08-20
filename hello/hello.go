package main

import "fmt"

// func main() {
// Side effect = printing to stdout
// fmt.Println("Hello, world") // Current domain -> the string we send t outside
// }

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
