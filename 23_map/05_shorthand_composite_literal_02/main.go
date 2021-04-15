package main

import "fmt"

func main() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Bat",
		"C": "Cat",
		"D": "Dog",
		"E": "Egg",
	}
	fmt.Println(myMap)
	fmt.Println(myMap["E"])

}
