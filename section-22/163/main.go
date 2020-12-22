package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	wg.Add(2)

	// send
	go foo(c)

	// receive
	go bar(c)

	fmt.Println("about to exit")
	wg.Wait()
}

func foo(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- i
	}
	wg.Done()
}

func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}
