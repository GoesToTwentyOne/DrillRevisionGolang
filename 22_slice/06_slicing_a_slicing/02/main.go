package main

import "fmt"

func main() {

	y := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println(y)
	fmt.Print("[2] ")
	fmt.Println(y[2])
	fmt.Print("[0:5] ")
	fmt.Println(y[:5])
	fmt.Print("[:] ")
	fmt.Println(y[:])
}
