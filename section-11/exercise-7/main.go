package main

import "fmt"

func main() {
	info := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hellooooo, James."},
	}
	fmt.Println(info)
	for i, record := range info {
		for j, datum := range record {
			fmt.Println(i, j, datum)
		}
	}
}
