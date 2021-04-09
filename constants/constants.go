package main

import "fmt"

func main() {

	//declaring
	const nam = "Tharun"
	const myConst int = 32
	fmt.Println(myConst, nam)

	const (
		name = "tharun"
		age  = "21"
	)
	fmt.Println(name, age)
	//untyped constants(not given type explicitly)
	const a = 3
	const f = 5.000
	//typed constants(type gien explicitly)
	var myInt int = 1
	var myFloat float64 = 1
	fmt.Println(myInt, myFloat)

}
