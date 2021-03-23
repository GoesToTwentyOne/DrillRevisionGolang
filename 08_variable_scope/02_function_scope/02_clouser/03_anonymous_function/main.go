package main

import (
	"fmt"
	"os"
)

func main() {
	x := 5
	ano := func() int {
		x++
		if x == 10 {
			os.Exit(1)
		}
		return x

	}
	fmt.Println(ano())
}
