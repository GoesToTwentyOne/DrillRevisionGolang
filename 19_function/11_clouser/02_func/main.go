package main

import "fmt"

var x int = 1 //package level

func main() {
	fmt.Println(x)

	fmt.Println(inc())
	fmt.Println(inc())
}
func inc() int {
	x++
	return x
}

//function in  function called clouser function
// closure helps us limit the scope of variables used by multiple functions
// without closure, for two or more funcs to have access to the same variable,
// that variable would need to be package scope
