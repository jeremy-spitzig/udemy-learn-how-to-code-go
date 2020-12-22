package main

import "fmt"

func main() {
	foo(func() {
		fmt.Println("This is inside an anonymous function.")
	})
}

func foo(f func()) {
	f()
}
