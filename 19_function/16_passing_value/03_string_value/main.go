package main

import "fmt"

func main() {
	n := "python"
	fmt.Println(n)
	changeME(n)
	fmt.Println(n)

}
func changeME(name string) {
	fmt.Println(name)
	name = "tiny go"
	fmt.Println(name)
}
