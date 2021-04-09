package main

import "fmt"

func main() {
	m1 := map[int]string{
		2: "hello",
		3: "hey",
	}
	m1[4] = "okk"
	m1[3] = "hey1"
	fmt.Println(m1)
	delete(m1, 2)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	val, ok := m1[3]
	fmt.Println(val, ok)
}
