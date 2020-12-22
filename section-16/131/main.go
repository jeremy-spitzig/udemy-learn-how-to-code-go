package main

import "fmt"

func main() {

	x := 0
	foo(x)
	fmt.Println(x)
	bar(&x)
	fmt.Println(x)
}

func foo(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
}

func bar(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	*x = 43
	fmt.Println(x)
	fmt.Println(*x)
}
