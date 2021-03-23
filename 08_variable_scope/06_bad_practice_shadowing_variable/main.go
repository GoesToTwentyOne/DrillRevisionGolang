package main

import "fmt"

func main() {
	add := add(8, 7) // add is now the result, not the function
	fmt.Println(add)

}
func add(a, b int) int {
	return a + b
}

// don't do this; bad coding practice to shadow variables
