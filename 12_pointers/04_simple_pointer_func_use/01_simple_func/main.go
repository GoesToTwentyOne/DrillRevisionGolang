package main

import "fmt"

func main() {
	x := 21
	goes(x)
	fmt.Println(x) //not change value

}
func goes(n int) {
	//n is over written not use

	// fmt.Println(n)
	n = 1

}
