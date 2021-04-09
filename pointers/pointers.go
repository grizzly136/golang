package main

import "fmt"

func main() {
	var p *int

	a := 2
	p = &a

	fmt.Println(a, p, &a, *p)
	*p = 3
	fmt.Println(a, p, &a, *p)
	//using new
	ptr := new(int)
	fmt.Println(ptr, *ptr)
	*ptr = 20
	fmt.Println(ptr, *ptr)
}
