package main

import "fmt"

type goes int

func main() {
	var x goes = 21
	fmt.Printf("%T    ", x)
	fmt.Println(x)

	var y int = 21 //statics type

	fmt.Printf("%T    ", y)
	fmt.Println(y)

	fmt.Println(int(x) + y)

	//don't work x+y

}
