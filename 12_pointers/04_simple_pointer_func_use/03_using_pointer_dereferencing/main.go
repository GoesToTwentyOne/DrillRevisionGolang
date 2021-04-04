package main

import "fmt"

func main() {
	x := 21
	fmt.Println(&x)
	fmt.Println(x)
	goes(&x)
	fmt.Println(x) // x is 5

}
func goes(g *int) {
	fmt.Println(&g)
	//dereferencing
	*g = 5

}
