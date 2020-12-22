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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Num go routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)
}
