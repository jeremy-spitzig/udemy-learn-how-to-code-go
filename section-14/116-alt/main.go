package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evens := filter(func(x int) bool { return x%2 == 0 }, ii...)
	s := reduce(0, func(acc int, val int) int { return acc + val }, evens...)
	fmt.Println(s)
}

func filter(pred func(int) bool, xi ...int) []int {
	var xj []int
	for _, x := range xi {
		if pred(x) {
			xj = append(xj, x)
		}
	}
	return xj
}

func reduce(init int, acc func(int, int) int, xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
