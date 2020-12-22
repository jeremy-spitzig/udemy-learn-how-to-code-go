package main

import "fmt"

func main() {
	defer func() { fmt.Println("This should happen last") }()
	fmt.Println("This should happen first")
	fmt.Println("This should happen second")
}
