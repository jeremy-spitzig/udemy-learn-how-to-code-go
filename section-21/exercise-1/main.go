package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	fmt.Println("I'm in foo!!!")
	wg.Done()
}

func bar() {
	fmt.Println("I'm in bar!!!")
	wg.Done()
}
