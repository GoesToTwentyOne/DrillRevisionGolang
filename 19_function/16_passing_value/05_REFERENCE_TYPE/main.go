package main

import "fmt"

func main() {
	v := make(map[string]int)
	changeME(v)
	fmt.Println(v["Nihad"])

}
func changeME(n map[string]int) {
	n["Nihad"] = 21
}
