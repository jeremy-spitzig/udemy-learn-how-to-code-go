package main

import (
	"fmt"
)

func main() {
	d := 10
	for i := 10; i <= 100; i++ {
		if i%d == 0 {
			fmt.Printf("%v is divisible by %v\n", i, d)
		} else if i%d == 1 || i%d == d-1 {
			fmt.Printf("%v is almost divisible by %v\n", i, d)
		} else {
			fmt.Printf("%v isn't even close to being divisible by %v\n", i, d)
		}
	}
}
