package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "James"
	(*p).last = "Bond"
}

func main() {
	p := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
