package main

import "fmt"

type square struct {
	side int
}

type circle struct {
	radius int
}

func (s square) area() float32 {
	return float32(s.side * s.side)
}

func (c circle) area() float32 {
	return 3.14 * float32(c.radius*c.radius)
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Printf("The area of the shape %f\n", s.area())
}

func main() {
	sq := square{
		side: 12,
	}
	circ := circle{
		radius: 12,
	}

	info(sq)
	info(circ)
}
