package main

import "fmt"

type goes int //create user define type

func main() {
	var x goes = 21
	fmt.Println(x)
	fmt.Printf("%T \n", x)

}
