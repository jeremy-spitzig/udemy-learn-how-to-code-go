package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type message struct {
	value   int
	channel int
}

func main() {
	channels := send()
	c := receive(channels...)
	for v := range c {
		fmt.Printf("%03d\t%v\n", v.channel, v.value)
	}
	fmt.Println("Done")
}

func send() []<-chan message {
	numMessages := rand.Intn(100)
	numChannels := rand.Intn(100)
	fmt.Printf("Sending %v messages on %v channels\n", numMessages, numChannels)
	channels := make([]<-chan message, numChannels)
	for i := 0; i < numChannels; i++ {
		c := make(chan message)
		channels[i] = c
		go func(channelNum int) {
			for i := 0; i < numMessages; i++ {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
				c <- message{
					rand.Int(),
					channelNum,
				}
			}
			close(c)
		}(i)
	}
	return channels
}

func receive(channels ...(<-chan message)) <-chan message {
	c := make(chan message)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))
		for i := 0; i < len(channels); i++ {
			channel := channels[i]
			go func() {
				for v := range channel {
					c <- v
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()
	return c
}
