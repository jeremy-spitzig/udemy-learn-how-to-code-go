package main

import (
	"fmt"
)

func main() {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Println(c)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", c)
		}
	}
}
