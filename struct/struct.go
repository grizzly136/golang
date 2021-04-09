package main

import "fmt"

type vehicle struct {
	Wheels int
	mode   string
}

func main() {
	v1 := vehicle{
		Wheels: 3,
		mode:   "fly",
	}
	ptr := &v1
	fmt.Println(v1, ptr, *ptr)
	fmt.Println(v1.Wheels)
	//initialization using new function
	v2 := new(vehicle)
	v2.Wheels = 8
	fmt.Println(v2)

}
