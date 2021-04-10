package main

import "fmt"

func main() {
	show(func(n int) {
		fmt.Println(n)

	}, []int{1, 2, 3, 4, 5})
}
func show(callback func(n int), numbers []int) {
	for _, value := range numbers {
		callback(value)

	}

}
