package main

import "fmt"

func main() {
	// when calling greeting, pass in an argument
	greeting("Nihad")
	greeting("Alex")

}

// greeting is declared with a parameter
func greeting(name string) {
	fmt.Println("Hi! ", name)

}
