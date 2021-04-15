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
	myMap["G"] = "Gan"
	fmt.Println(myMap)
	fmt.Println(myMap["G"])
	fmt.Println(len(myMap))

}
