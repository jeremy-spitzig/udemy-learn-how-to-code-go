package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {
	t1 := struct {
		vehicle
		fourWheel bool
	}{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}
	s1 := struct {
		vehicle
		luxury bool
	}{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.fourWheel)
	fmt.Println(s1.luxury)
}
