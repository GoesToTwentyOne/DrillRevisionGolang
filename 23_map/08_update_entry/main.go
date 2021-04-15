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
	myMap["F"] = "Fan"
	fmt.Println(myMap)
	myMap["F"] = "Forg"
	fmt.Println(myMap)
	myMap["C"] = "Cartoon"
	fmt.Println(myMap)

}
