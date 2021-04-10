package main

import "fmt"

func main() {

	// A Higher-Order function is a function that receives a function as an argument or returns
	//the function as output.
	// Higher order functions are functions that operate on other functions, either by taking them as
	//arguments or by returning them.

	//callback func
	//A "callback func" is when we pass a func into a func as  an argument.
	//pass a func into a func as an argument

	//https://www.golangprograms.com/
	//Tood sir callback func video
	//https://www.youtube.com/watch?v=Qsp9htNRRE8&t=361s
	v := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return 1
		}
		return xi[0] + xi[len(xi)-1]

	}
	g := foo(v, []int{1, 2, 3, 4, 5})
	fmt.Println(g)
}
func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n

}
