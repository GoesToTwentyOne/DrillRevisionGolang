package main

import "fmt"

func main() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Bat",
		"C": "Cat",
		"D": "Dog",
	}
	fmt.Println(myMap)
	delete(myMap, "B")
	fmt.Println(myMap)
}
