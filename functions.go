package main

import "fmt"

func add(values ...int) (res int, avg int) {
	res = 0
	for _, v := range values {
		res += v
	}
	avg = res / len(values)
	return
}

func div(v1, v2 int) (int, error) {
	if v2 == 0 {
		return 0, fmt.Errorf("invalid")
	} else {
		return v1 / v2, nil
	}

}

func main() {
	a, av := add(1, 2, 3, 4, 5)
	fmt.Println(a, av)
	fmt.Println(div(4, 2))
}
