package main

import "fmt"

type numbers struct {
	num1 int
	num2 int
}

func main() {
	n1 := numbers{
		num1: 4,
		num2: 5,
	}
	fmt.Println(n1.isEven())
	fmt.Println(n1.isEven())

}
func (n *numbers) isEven() bool {
	if (n.num1)%2 == 0 && (n.num2)%2 == 0 {
		return true
	} else {
		n.num1 = 2
		n.num2 = 4
		return false
	}
}
