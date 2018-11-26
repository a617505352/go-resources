package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var on sync.Once

func setup() {
	fmt.Println("Init")
}

func doStuff() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	go doStuff()
	go doStuff()
	wg.Wait()
}
