package main

import "fmt"

func main() {
	var aChannel chan int
	fmt.Println("value ", aChannel)
	fmt.Printf("Type %T \n", aChannel)

	bChannel := make(chan int)
	fmt.Println("value ", bChannel)
	fmt.Printf("Type %T \n", bChannel)
}
