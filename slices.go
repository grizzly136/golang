package main

import "fmt"

func main() {
	//slices
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a := slc[2:7]
	fmt.Println(slc)
	a[1] = 12
	fmt.Println(a)
	//slice using make
	s1 := make([]int, 4, 10)
	s1[2] = 5
	fmt.Println(s1)
	//spreading
	s2 := append(a, s1...)
	fmt.Println(s2)

}
