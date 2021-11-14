package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	// Receive ONLY Channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// Send ONLY channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
