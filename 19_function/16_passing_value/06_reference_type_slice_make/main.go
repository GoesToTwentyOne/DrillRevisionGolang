package main

import "fmt"

func main() {
	v := make([]string, 1, 25)
	chaneMe(v)
	fmt.Println(v)
	fmt.Println(cap(v))

}
func chaneMe(n []string) {
	n[0] = "Nihad"

}
