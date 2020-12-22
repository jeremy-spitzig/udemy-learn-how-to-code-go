package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	f1 := bar()
	fmt.Printf("%T\n", f1)
	v := f1()
	fmt.Println(v)
}

func foo() string {
	s := "Hello, World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
