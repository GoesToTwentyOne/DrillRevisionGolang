package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(avarage(data))

}
func avarage(n []float64) float64 {
	var sum float64 = 0
	for _, v := range n {
		sum += v
	}
	return sum / float64(len(n))

}
