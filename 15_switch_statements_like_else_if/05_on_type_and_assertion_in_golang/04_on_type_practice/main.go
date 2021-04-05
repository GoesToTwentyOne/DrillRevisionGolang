package main

import "fmt"

func main() {
	termiter(77)
	termiter(true)
	termiter('N')
	termiter(25.2)

	//the struct type

	var acc = account{"Md.Nihad Hossain", 19878888, "December"}
	termiter(acc)
	termiter(acc.Name)
	termiter(acc.Ac)
	termiter(acc.DateToOpenMonth)

}

type account struct {
	Name            string
	Ac              int
	DateToOpenMonth string
}

func termiter(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer type")
	case string:
		fmt.Println("String type")
	case account:
		fmt.Println("struct account type")
	case bool:
		fmt.Println("bool type")
	case float64:
		fmt.Println("float type")
	case rune:
		fmt.Println("rune type")
	default:
		fmt.Println("Don'y miss the useful type")
	}

}
