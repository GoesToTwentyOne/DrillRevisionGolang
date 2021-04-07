package main

import "fmt"

func main() {
	fmt.Println(fullname("Nihad ", "Hossain"))

}
func fullname(fname, lname string) (string, string) {
	return fname, lname

}
