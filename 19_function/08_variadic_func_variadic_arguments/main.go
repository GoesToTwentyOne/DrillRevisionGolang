package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 10, 25, 8, 7, 7.8, 4.72}

	fmt.Println(avarage(data...))

}

func avarage(n ...float64) float64 {
	var sum float64 = 0
	for _, v := range n {
		sum += v

	}
	return sum / float64(len(n)) //type conversion

}
