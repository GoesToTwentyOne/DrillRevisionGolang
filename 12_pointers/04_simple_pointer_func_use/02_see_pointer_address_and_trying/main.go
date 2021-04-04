package main

import "fmt"

func main() {
	x := 21
	fmt.Println("func main : ", &x)
	goes(x)
	fmt.Println(x)

}
func goes(g int) {
	fmt.Printf("func goes 01 %p \n", &g)
	fmt.Printf("func goes 02 %v \n", &g)
	fmt.Printf("func goes 03 %b \n", &g)
	g = 0

}
