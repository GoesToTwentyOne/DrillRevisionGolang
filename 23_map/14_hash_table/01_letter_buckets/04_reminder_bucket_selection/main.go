package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, ".....", string(rune(i)), "......", i%12)
	}
}
