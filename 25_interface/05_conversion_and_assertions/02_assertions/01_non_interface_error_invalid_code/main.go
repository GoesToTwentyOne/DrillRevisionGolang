package main

import "fmt"

func main() {
	name := "Sydney"
	// //str, ok := name.(string)
	// if ok {
	// 	fmt.Printf("%q\n", str)
	// } else {
	// 	fmt.Printf("value is not a string\n")
	// }
	fmt.Println(name) //only interface type allow assertion

}
