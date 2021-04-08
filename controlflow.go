package main

import "fmt"

func main() {
	var x = 24
	//if
	if x > 25 {
		fmt.Println("greater than 25")
	} else if x < 25 {
		fmt.Println("less than 25")
	} else {
		fmt.Println("its 25")
	}
	//switch

	switch c := 2; c {
	case 1:
		fmt.Println("its  1")
	case 2, 4, 6:
		fmt.Println("its  2,4,6")
	case 3:
		fmt.Println("its  3")
	default:
		fmt.Println("none")
	}
	//for
	for i := 0; i < 5; i++ {

		if i == 1 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

}
