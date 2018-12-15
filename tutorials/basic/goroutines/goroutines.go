/*
	Goroutines
		- Goroutines are functions or methods that run concurrently with other functions or methods.
		- Goroutines can be thought of as light weight threads.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	basic()
	atomicFunctions()
	mutexes()
}

/*
====== BASIC ======
*/
// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func basic() {
	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(1)
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s\n", prefix)
	}
	fmt.Println("Completed", prefix)
}

// ===== Atomic Functions =====
func atomicFunctions() {

}

// ===== Mutexes =====
var (
	counter int
	// wg is used to wait for the program to finish.
	wgM sync.WaitGroup
	// mutex is used to define a critical section of code.
	mutex sync.Mutex
)

func mutexes() {
	wgM.Add(2)
	go incCounter(1)
	go incCounter(2)
	wgM.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wgM.Done()

	mutex.Lock()
	// critical section
	{
		value := counter
		// Yield the thread and be placed back in queue.
		// the scheduler assigns the same goroutine to continue running
		runtime.Gosched()
		value++
		counter = value
	}
	mutex.Unlock()
}
