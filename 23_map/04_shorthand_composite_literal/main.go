package main

import "fmt"

func main() {
	myMap := map[string]string{}
	myMap["A"] = "Apple"
	myMap["B"] = "Bat"
	myMap["C"] = "Cat"
	myMap["D"] = "Dog"
	fmt.Println(myMap)
	fmt.Println(myMap["D"])
}
