package main

import "fmt"

func main() {
	if true && false { //don't execute this block
		fmt.Println("I'm not run")
	}
	fmt.Println("and operation")
}
