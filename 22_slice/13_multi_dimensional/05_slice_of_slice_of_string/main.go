package main

import "fmt"

func main() {
	var records [][]string

	//student1 data
	student1 := make([]string, 4)
	student1[0] = "Nihad"
	student1[1] = "Rowjatul"
	student1[2] = "Neha"
	student1[3] = "Alex"
	records = append(records, student1)
	fmt.Println(records)
	//student2 data
	student2 := make([]string, 4)

	student2[0] = "1"
	student2[1] = "2"
	student2[2] = "3"
	student2[3] = "4"
	records = append(records, student2)
	fmt.Println(records)
	//student3 data
	student3 := make([]string, 4)

	student3[0] = "J"
	student3[1] = "k"
	student3[2] = "l"
	student3[3] = "m"
	records = append(records, student3)
	fmt.Println(records)
	//student4 data
	student4 := make([]string, 4)
	student4[0] = "88"
	student4[1] = "99"
	student4[2] = "111"
	student4[3] = "222"
	records = append(records, student4)
	fmt.Println(records)

}
