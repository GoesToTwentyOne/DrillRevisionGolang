package main

import "fmt"

func main() {
	var x int32 = 'A'
	var y rune = 'B'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))

}

// rune is an alias for int32; normally omitted in this statement
