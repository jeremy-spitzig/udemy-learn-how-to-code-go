package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	fmt.Println(x)
	for i := len(x); i < cap(x); i++ {
		x = append(x, i)
	}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)
	x = append(x, 100)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)
}
