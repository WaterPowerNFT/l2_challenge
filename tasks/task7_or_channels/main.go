package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{}, 1)
	var wg sync.WaitGroup
	closeAllChannel := make(chan bool, 1)
	defer close(closeAllChannel)
	wg.Add(len(channels))
	for _, elem := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-closeAllChannel:
				closeAllChannel <- true
			case <-ch:
				closeAllChannel <- true
				out <- 1
			}
			wg.Done()

		}(elem)
	}
	wg.Wait()
	return out
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(99*time.Millisecond),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
