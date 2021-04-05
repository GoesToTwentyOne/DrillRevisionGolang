package main

import "fmt"

func main() {
	switch "Nihad" {
	case "Tonay", "Tahmid":
		fmt.Println("This is Tonay or Tahmid")
	case "Nihad", "Rowjatul":
		fmt.Println("This is Nihad or Rowjatul")
	case "Mahi", "Mithila":
		fmt.Println("This is Mahi or Mithila")
	case "Alex", "Anna":
		fmt.Println("This is Alex or Anna")
	default:
		fmt.Println("This don't match")
	}
}
