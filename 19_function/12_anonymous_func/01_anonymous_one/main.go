package main

import "fmt"

func main() {
	func(x, y int) {
		fmt.Println(x * y)

	}(10, 20)

}

/*Go language provides a special feature known as an anonymous function.
An anonymous function is a function which doesn’t contain any name. It is useful when you want to create
an inline function. In Go language,
an anonymous function can form a closure. An anonymous function is also known as function literal.*/
