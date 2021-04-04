package main

import "fmt"

func main() {
	x := 21
	fmt.Println(&x)
	goes(&x)
	fmt.Println(&x)
	fmt.Println(x)

}
func goes(g *int) {
	fmt.Println(g)
	*g = 5

}
