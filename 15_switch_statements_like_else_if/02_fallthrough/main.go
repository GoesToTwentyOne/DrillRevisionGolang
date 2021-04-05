package main

import "fmt"

func main() {
	// With the help of fallthrough statement, we can use to transfer the program control
	// just after the statement is executed in the switch cases even if the expression does not match.
	// Normally, control will come out of the statement of switch just after the execution of first line after
	// match.
	//  Don’t put the fallthrough in the last statement of switch case.
	switch "Neha" {
	case "Nihad":
		fmt.Println("This is Nihad Hossain")

	case "Rowjatul":
		fmt.Println("This is Rowjatul")
	case "Neha":
		fmt.Println("This is Neha")
		fallthrough
	case "Shova":
		fmt.Println("This is Shova")
	default:
		fmt.Println("Don't match ")
	}
}
