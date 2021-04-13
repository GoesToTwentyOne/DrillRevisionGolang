package main

import "fmt"

func main() {
	n := "Python"
	fmt.Println(&n)
	fmt.Println(n)
	changeME(&n)
	fmt.Println(&n)
	fmt.Println(n)

}
func changeME(name *string) {
	fmt.Println(name)
	fmt.Println(*name)
	*name = "Tiny go"
	fmt.Println(name)
	fmt.Println(*name)
}
