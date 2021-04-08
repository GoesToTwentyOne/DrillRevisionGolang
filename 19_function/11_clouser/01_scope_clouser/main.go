package main

import "fmt"

func main() {
	x := 21
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "be positive always"
		fmt.Println(y)
	}
	//fmt.Println(y)  it isn't y scope

}
