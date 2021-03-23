package main

import (
	"fmt"
	"os"
)

func main() {
	inc := wrapping()
	fmt.Println(inc())

}
func wrapping() func() int {
	x := 0
	return func() int {
		x++
		if x == 5 {
			os.Exit(1)
		}
		return x

	}
}
