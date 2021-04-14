package main

import "fmt"

func main() {
	var smallAlphabet [26]string
	for i := 97; i <= 122; i++ {
		smallAlphabet[i-97] = string(i)
	}
	fmt.Println(smallAlphabet)
	fmt.Println(smallAlphabet[25])

}
