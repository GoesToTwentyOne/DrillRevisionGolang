package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("HI! my name is Nihad")
	}
	greeting()
	fmt.Printf("%T \n", greeting)
}
