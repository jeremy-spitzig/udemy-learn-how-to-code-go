package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("The inner loop: %d\n", j)
		}
	}

	for pos, char := range "Hello, World." {
		fmt.Printf("The character at position %d is %c.\n", pos, char)
	}
}
