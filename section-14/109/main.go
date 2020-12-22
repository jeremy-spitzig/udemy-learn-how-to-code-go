package main

import "fmt"

func main() {
	sum := foo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
