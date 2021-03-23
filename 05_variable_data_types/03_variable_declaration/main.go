package main

import "fmt"

func main() {
	var name string        //variable declaration
	name = "Nihad Hossain" //value assign
	fmt.Println(name)
	//more than one at once
	var a, b, c, d int
	a = 5
	fmt.Println(a, b, c, d)
	//string more than one at once assign value
	var e, f, g string = "Nihad", "Shova", "Mim"
	fmt.Println(e, f, g)
	//Type inference
	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	//When the right hand side of the declaration is typed, the new variable is of that same type:

	//infer type
	var name2 = "Nihad"
	fmt.Println(name2)
	//infer mixed up types
	var j, i, l = 4, false, 'N'
	fmt.Printf("%v  %v  %v %T %T %T", j, i, l, j, i, l)

}
