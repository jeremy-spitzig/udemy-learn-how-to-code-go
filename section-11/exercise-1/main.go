package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)
}
