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
	c <- 42
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
