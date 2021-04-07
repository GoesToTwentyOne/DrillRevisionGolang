package main

import "fmt"

func main() {
	fmt.Println(greeting("Nihad ", "Hossain"))

}
func greeting(fname, sname string) (fullname string) {
	fullname = fname + sname
	return
	//Avoid using named returns.

}
