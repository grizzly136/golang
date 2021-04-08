package main

import "fmt"

//initialiation in package level
var m int = 32
var (
	l int = 1
	n int = 3
)

func main() {
	// initailization in function level

	var i int
	i = 5
	var j int = 6
	k := 7

	fmt.Println(i, j, k)

	//conversion
	var m int = 1
	var n float32

	n = float32(m)
	p := string(m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", p, p)

	//primitives
	//bool(false by default)
	var b bool
	c := 2 == 2
	fmt.Println(b, c)

	//numeric

	//integers
	//float
	//complex
	//string

}
