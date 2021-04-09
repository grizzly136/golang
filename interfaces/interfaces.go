package main

import "fmt"

type feelings interface {
	happy() int
	sad() int
	confused() int
}
type person struct {
	name string
	age  int
}

func main() {

	p1 := person{
		name: "myself",
		age:  18,
	}

	fmt.Println(p1.happy(), p1.sad())
}

func (p person) happy() bool {
	if p.age < 18 {
		return true
	} else {
		return false
	}
}

func (p person) sad() bool {
	if p.age > 18 {
		return true
	} else {
		return false
	}
}
func feel(f feelings) {
	fmt.Println(f.happy())
	fmt.Println(f.sad())
}
