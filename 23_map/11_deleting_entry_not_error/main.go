package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Nihad",
		1: "Alex",
		2: "Anna",
		3: "Candy",
	}
	fmt.Println(myMap)
	delete(myMap, 5) //don't return error
	fmt.Println(myMap)

}
