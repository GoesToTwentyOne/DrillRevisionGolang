package main

import "fmt"

func main() {
	func(x, y int) {
		fmt.Println("Result : ", x+y)

	}(4, 5)
}
