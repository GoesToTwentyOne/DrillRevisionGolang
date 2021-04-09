package main

import "fmt"

func main() {
	value := sumeven()
	v2 := value([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(v2)

}
func sumeven() func(n []int) int {
	v := func(n []int) int {
		var sum int = 0
		for i := 0; i < len(n); i++ {
			if n[i]%2 == 0 {
				sum += n[i]
			}

		}
		return sum

	}
	return v

}
