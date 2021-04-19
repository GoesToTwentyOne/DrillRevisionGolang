package main

import (
	"fmt"
)

func main() {
	foo()

	bar()

}
func foo() {
	for i := 0; i < 10; i++ {

		// time.Sleep(time.Second)
		fmt.Println("foo  : ", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		// time.Sleep(time.Second)
		fmt.Println("bar  : ", i)
	}
}
