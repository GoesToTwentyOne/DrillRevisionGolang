package main

import "fmt"

func main() {
	//setup the pipelines
	intTOch := gen(4, 5, 6, 7)
	ch := sq(intTOch)
	for values := range ch {
		fmt.Println(values)
	}

}

func gen(nums ...int) chan int {
	chOut := make(chan int)
	go func() {
		for _, value := range nums {
			chOut <- value

		}
		close(chOut)
	}()
	return chOut

}
func sq(input chan int) chan int {
	chOut := make(chan int)
	go func() {
		for values := range input {
			chOut <- values * values

		}
		close(chOut)
	}()
	return chOut

}

//golang-pipelines:  t.ly/byTG
