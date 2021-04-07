package main

import "fmt"

func main() {
	avg(5, 10, 20)

}
func avg(n ...float64) {
	fmt.Println(n)
	fmt.Printf("%T \n", n) //float64 array type

	//variadic function check
	fmt.Println(avarage(1, 2, 3, 4, 5, 10, 25, 8, 7, 7.8, 4.72))
}
func avarage(n ...float64) float64 {
	var sum float64 = 0
	for _, v := range n {
		sum += v

	}
	return sum / float64(len(n)) //type conversion

}
