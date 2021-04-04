package main

import "fmt"

func main() {
	for i := 1; i < 90; i++ {
		if i%2 == 0 {
			fmt.Println("even ", i)
		} else {
			fmt.Println("odd ", i)
		}
	}
}
