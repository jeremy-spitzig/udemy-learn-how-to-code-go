package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := foo(xi...)
	s2 := bar(xi)
	fmt.Println(s1)
	fmt.Println(s2)
}

func foo(xi ...int) int {
	return bar(xi)
}

func bar(xi []int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
