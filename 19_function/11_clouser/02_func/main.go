package main

import "fmt"

var x int = 1 //package level

func main() {
	fmt.Println(x)

	fmt.Println(inc())
	fmt.Println(inc())
}
func inc() int {
	x++
	return x
}
