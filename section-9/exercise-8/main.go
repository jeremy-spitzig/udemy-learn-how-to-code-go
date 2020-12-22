package main

import (
	"fmt"
)

func main() {
	switch {
	case 10 < 5:
		fmt.Println("10 < 5")
	case 10 <= 5:
		fmt.Println("10 <= 5")
	case 10 > 5:
		fmt.Println("10 > 5")
	case 10 >= 5:
		fmt.Println("10 >= 5")
	}
}
