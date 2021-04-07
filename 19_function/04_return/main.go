package main

import "fmt"

func main() {

	fmt.Println(greeting("Nihad", "Hossain"))

}
func greeting(fname, sname string) string {
	return fmt.Sprintf("HI! %s %s", fname, sname)

}
