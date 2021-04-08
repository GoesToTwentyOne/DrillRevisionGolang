package main

import "fmt"

func main() {
	f := greeting()
	fmt.Println(f())
	fmt.Printf("%T \n", f)

}

//              return a function
func greeting() func() string {
	return func() string {
		return "Hello wold"

	}

}
