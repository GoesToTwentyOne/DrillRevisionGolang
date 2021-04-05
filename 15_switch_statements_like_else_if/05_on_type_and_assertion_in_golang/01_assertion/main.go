package main

import "fmt"

func main() {

	//1. Type assertions in Golang provide access to the exact type of variable of an interface.
	// If already the data type is present in the interface, then it will retrieve the actual data type
	// value held by the interface. A type assertion takes an interface value and extracts from it a value
	// of the specified explicit type. Basically, it is used to remove the ambiguity from the interface variables.
	//2. To test whether an interface value holds a specific type, a type assertion can return
	// two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	//3.If i holds a s, then t will be the underlying value and ok will be true.
	var i interface{} = "Nihad"
	s, ok := i.(string)
	fmt.Println(s, ok)
	//If not, ok will be false and t will be the zero value of type s2, and no panic occurs.
	s2, unsuccessful := i.(int)
	fmt.Println(s2, unsuccessful)

}
