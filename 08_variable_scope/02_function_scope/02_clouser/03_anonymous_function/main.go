package main

import (
	"fmt"
	"os"
)

func main() {
	x := 5
	ano := func() int {
		x++
		if x == 10 {
			os.Exit(1)
		}
		return x

	}
	fmt.Println(ano())
	/*closure helps us limit the scope of variables used by multiple functions
	without closure, for two or more funcs to have access to the same variable,
	that variable would need to be package scope*/

	//anonymous function
	//a function without a name
	//syntax
	//func expression
	//assigning a func to a variable
}
