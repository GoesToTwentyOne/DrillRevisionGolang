package main

import "fmt"

func main() {
	v := func(x int) {
		fmt.Println(x)

	}
	goes(v, []int{1, 2, 3})

}
func goes(callback func(x int), numbers []int) {
	for _, value := range numbers {
		callback(value)
	}

}
