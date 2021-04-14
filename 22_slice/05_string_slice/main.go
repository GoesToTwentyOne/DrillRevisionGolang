package main

import "fmt"

func main() {
	myFamily := []string{
		"Mother",
		"Father",
		"sister",
		"ME",
	}
	for i, value := range myFamily {
		fmt.Println(i, value)
	}

	// for j := 0; j < len(myFamily); j++ {
	// 	fmt.Println(myFamily[j])
	// }

}
