package main

import (
	"fmt"

	vargo "github.com/GoesToTwentyOne/DrillRevisionGolang/08_variable_scope/03_package_level_visiability/able"
)

func main() {
	fmt.Println(vargo.FirstName)
	vargo.Varprint()

}
