package main

import (
	"fmt"
)

func main() {
	year := 1985
	for {
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}
