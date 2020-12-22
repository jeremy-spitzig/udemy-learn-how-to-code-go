package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Printf("I am %s %s\n", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"James", "Bond"}
	saySomething(&p)
	p.speak()
	// can't actually treat p as a human: saySomething(p)
}
