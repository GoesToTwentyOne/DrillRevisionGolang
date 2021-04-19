package main

import "fmt"

func main() {
	var x int = 450
	var y float64 = 5.5
	//fmt.Println(x+y)///miss mached type
	fmt.Println(x + int(y))
}
