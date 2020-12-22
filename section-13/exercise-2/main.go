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

	people := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range people {
		fmt.Println(p.first)
		fmt.Println(p.last)
		for i, v := range p.ficf {
			fmt.Println(i, v)
		}
	}
}
