package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
	fmt.Scanln()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Second)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Second)
	}
}
