package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)
	rec(eve, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func rec(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			}
			fmt.Println("from comma ok:", i)
		}
	}
}
