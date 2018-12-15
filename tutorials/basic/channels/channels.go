/*
	Channels
	- The way Goroutines communicate.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbufferedChannels()
	bufferedChannels()
}

/*
	Unbuffered channels
		-
*/
var wg sync.WaitGroup

func unbufferedChannels() {
	// Create an unbuffered channel
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		close(baton)
		return
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}

/*
	Buffered channels
	 	- A buffered channel is a channel with capacity to hold one or more values before they’re received.
*/
func bufferedChannels() {

}
