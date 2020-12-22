package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
