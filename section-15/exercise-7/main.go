package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is inside an anonymous function.")
	}
	f()
}
