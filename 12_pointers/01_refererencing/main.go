package main

import "fmt"

func main() {
	a := 21
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	fmt.Println(b)
	fmt.Println(&b)
	// fmt.Println(*b)

}
