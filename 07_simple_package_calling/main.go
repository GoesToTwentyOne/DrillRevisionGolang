package main

import (
	"fmt"

	"github.com/GoesToTwentyOne/DrillRevisionGolang/07_simple_package_calling/add"
	"github.com/GoesToTwentyOne/DrillRevisionGolang/07_simple_package_calling/name"
	Testsum "github.com/GoesToTwentyOne/DrillRevisionGolang/07_simple_package_calling/sum_1_to_10"
)

func main() {
	fmt.Println(add.Add(20, 30))
	fmt.Println(Testsum.Sum())
	fmt.Println(name.Myorginal("Md.Nihad Hossain"))

}
