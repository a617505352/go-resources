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
	selectChannels()
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

/*
	Selct channels
	 	- A buffered channel is a channel with capacity to hold one or more values before they’re received.
*/
func selectChannels() {
	c1 := make(chan string)
	c2 := make(chan string)
	signal := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		signal <- true
	}()

loop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received %s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("received %s\n", msg2)
		case <-signal:
			fmt.Println("end process")
			break loop
		}
	}
}
