package main

import "fmt"

func main() {
	var assert interface{} = "My nane is Nihad Hossain"
	str1, ok := assert.(string)
	fmt.Println(str1, ok)
	str2, unsuccessful := assert.(int)
	fmt.Println(str2, unsuccessful)
	//return 0 and false
	str3 := assert.(int)
	fmt.Println(str3)
	// return panic error

}
