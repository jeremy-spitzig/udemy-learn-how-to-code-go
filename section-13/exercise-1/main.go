package main

import "fmt"

type person struct {
	first string
	last  string
	ficf  []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		ficf:  []string{"Vanilla", "Strawberry"},
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		ficf:  []string{"Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.ficf {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.ficf {
		fmt.Println(i, v)
	}
}
