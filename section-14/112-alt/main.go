package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first, h.(secretAgent).last)
	}
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	bar(sa1)
	bar(sa2)
	bar(p1)
}
