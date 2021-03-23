package main

import (
	"fmt"
	"os"
)

var x = 1

func main() {
	for {
		fmt.Println(inc())
	}

}
func inc() int {
	x++

	if x == 5 {
		os.Exit(1)

	}
	return x

	//closure helps us limit the scope of variables used by multiple functions
	//without closure, for two or more funcs to have access to the same variable,
	//that variable would need to be package scope

}
