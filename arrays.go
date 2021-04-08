package main

import "fmt"

func main() {
	//arrays
	a := [...]int{1, 2, 3}
	for k, v := range a {
		fmt.Println(k, v)
	}
	a[2] = 5
	for k, v := range a {
		fmt.Println(k, v)
	}
	//multidimensional
	matrix := [3][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9}}

	fmt.Println(matrix)

}
