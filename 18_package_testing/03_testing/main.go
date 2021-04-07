package main

import "fmt"

func main() {
	fmt.Println(IntMin(2, -2))
	fmt.Println(IntMinTableDriven())

}
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func IntMinTableDriven() bool {
	var structTest = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range structTest {
		// testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		ans := IntMin(tt.a, tt.b)
		if ans != tt.want {
			// fmt.Println("false")
			return false

		}
		//fmt.Println("true")

	}
	return true

}
