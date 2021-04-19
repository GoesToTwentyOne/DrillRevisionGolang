package main

import "fmt"

func main() {
	var x float64 = 450.89
	var y int = 5
	//fmt.Println(x+y)///miss mached type
	fmt.Println(x + float64(y))
}
