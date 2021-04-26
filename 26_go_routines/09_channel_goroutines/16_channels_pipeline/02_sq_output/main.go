package main

import "fmt"

func main() {
	//setup the pipelines
	for values := range sq(sq(gen(2, 3, 4, 5))) {
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
