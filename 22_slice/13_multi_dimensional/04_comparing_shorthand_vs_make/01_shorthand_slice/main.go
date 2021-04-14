package main

import "fmt"

func main() {
	x := []string{}
	y := [][]string{}
	// x[1] = "Nihad" //panic error
	x = append(x, "Nihad")
	fmt.Println(x)
	fmt.Println(y)

}
