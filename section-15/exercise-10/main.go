package main

import "fmt"

func main() {
	i := incrementor()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}

func incrementor() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
