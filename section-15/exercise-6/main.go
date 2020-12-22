package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is inside an anonymous function.")
	}()
}
