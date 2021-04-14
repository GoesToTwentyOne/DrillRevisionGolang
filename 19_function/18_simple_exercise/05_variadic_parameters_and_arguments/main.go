package main

import "fmt"

func main() {
	variadic(1, 2)
	variadic(1, 2, 5, 7, 8, 10)
	aslice := []int{1, 2, 3, 4, 5}
	variadic(aslice...)

}
func variadic(n ...int) {
	fmt.Println(n)
}
