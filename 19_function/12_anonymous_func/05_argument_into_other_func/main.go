package main

import "fmt"

func goes(i func(fname, lname string) string) {
	fmt.Println(i("Nihad", "Hossain"))
	fmt.Printf("%T \n", i)

}
func main() {
	v := func(fname, lname string) string {
		return fname + lname + "fun"

	}
	goes(v)

}
