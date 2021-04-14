package main

import "fmt"

func main() {
	x := make([]string, 21)
	y := make([][]string, 21)
	// x[0] = "Nihad"
	x = append(x, "Hossain") //add after capacity or langth
	fmt.Println(x)
	fmt.Println(y)
}
