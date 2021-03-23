package main

import "fmt"

func main() {
	x := 21
	fmt.Println(x)

}
func goes() {
	//dont access it another function
	//fmt.Println(x)
	x := 42
	fmt.Println(x)
}
