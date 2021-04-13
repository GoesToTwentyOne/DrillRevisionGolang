package main

import "fmt"

func main() {
	age := 21
	changeMe(age) //pass an argument
	fmt.Println(age)

}
func changeMe(n int) { //one parameter
	fmt.Println(n)
}
