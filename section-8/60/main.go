package main

import (
	"fmt"
)

func main() {
	var c byte
	for c = 33; c <= 122; c++ {
		fmt.Printf("%v\t%#04x\t%#U\t%c\n", c, c, c, c)
	}
}
