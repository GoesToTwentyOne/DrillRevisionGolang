package main

import "fmt"

func main() {
	myMap := map[int]string{
		01: "Nihad",
		02: "Rowjatul",
		03: "Neha",
		04: "Suba",
		05: "Ana",
		06: "Alex",
	}
	for key, value := range myMap {
		fmt.Println(key, "........", value)
	}
}
