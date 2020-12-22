package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	var wg sync.WaitGroup
	const gs = 1000
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
