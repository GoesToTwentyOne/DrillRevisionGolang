package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")

	}()

}

//Anonymous functions are useful when you want to define a function inline without having to name it.
