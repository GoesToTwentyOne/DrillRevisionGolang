package main

import "fmt"

func main() {
	a := 21

	fmt.Println(a)  // 21
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 21 dereferencing
	//* is an operator that case

}
