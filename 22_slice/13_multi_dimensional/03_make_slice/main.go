package main

import "fmt"

func main() {
	student := make([]string, 21)
	students := make([][]string, 21)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
