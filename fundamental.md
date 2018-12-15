<h1 align="center">Golang Fundamental</h1>

- [Goroutines and Channels](#goroutines-and-channels)
- [Concurrency with Shared Variables](#concurrency-with-shared-variables)
- [Testing](#testing)
- [Reflection](#reflection)
- [Others](#others)

# Goroutines and Channels

## Goroutines

- Like a thread in Go
- Many Goroutines executes within a single OS thread

![screen shot 2018-11-24 at 22 26 14](https://user-images.githubusercontent.com/11765228/48969328-271f5600-f038-11e8-985e-1b67e59c7667.png)

## Go Runtime Scheduler

- Schedules goroutines inside an OS thread
- Like a little OS indie a single OS thread
- Logical processor is mapped to a thread

![screen shot 2018-11-24 at 22 26 51](https://user-images.githubusercontent.com/11765228/48969327-271f5600-f038-11e8-99be-cbf9138b4451.png)

## Go Stateme

A go stateme causes the function to be called in a newly created goroutine. The go statement itself completes immediately:

```go
f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

## Exiting a Goroutin

- A goroutine exits when its code is complete
- When the main gorount is complete, all others goroutines exit

### WaitGroup

- Sync package contains functions to synchronize between goroutine
- sync.WaitGroup force a goroutine to wait for other goroutines
- Contains an internal counter

```go
wg.Add()  // increments the counter
wg.Done() // decrements the counter
wg.Wait() // blocks until counter === 0
```

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func process(i int, wg *sync.WaitGroup) {
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}

func main() {
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}
```

**It is important to pass the address of wg in line no. 21. If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.**

## Channels

- Transfer data between goroutines
- Use `make()` to creat a channel

```go
ch <- x // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch // a receive statement; result is discarded
```

### Unbuffered Channels

- Unbuffered channels cannot hold data in transit
- Sending blocks until data is received
- Receiving blocks until data is sent

```go
ch := make(chan type)
```

### Buffered Channels

- Channels can contain a limited number of objects
- Capacity is the number of objects it can hold in transit
- Optional argument to make()defines channel capacity `c := make(chan int, 3)`
- Sending only blocks if buffer is full
- Receiving only blocks if buffer is empty

```go
ch := make(chan type, capacity)
```

## Iterating Through a Channel

- Common to iteratively read from a channel

```go
for i := range c {
    fmt.Println(i)
}
```

- Continues to read from channel c
- One iteration each time a new value is received
- Iterates when sender calls `close(c)`

## Receiving from Multiple Goroutines

### Select Statement

- May have a choice of which data to use
- Use the select statement to wait on the first data from a set of channels

```go
select {
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
    default:
        fmt.Println("nop")
}
```

### Select with an Abort Channel

- Use select with a separate abort channel
- May want to receive data until an abort signal is received

```go
for {
    select {
        case a <- c:
            fmt.Println(a)
        case <- abort:
            return
    }
}
```

[[↑] Back to top](#golang-fundamental)

# Concurrency with Shared Variables

- Sharing variables concurrently can cause problems
- Tow goroutines writing to a shared variable can interfere with each other

## sync.Mutex

- A Mutex ensures mutual exclusion
- Vues a binary semaphore

```go
var i int = 0
var mut sync.Mutex

func inc() {
    mutex.Lock()
    x = x + 1
    mutex.Unlock()
}
```

## sync.Once

- Function f is executed only one time
- All calls to once.DO9() block until the first returns

```go
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
```

##

[[↑] Back to top](#golang-fundamental)

# Testing

## _Test_ Functions

```go
func TestName(t *testing.T) {
    // ...
}
```

## _Benchmark_ Functions

```go
func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}
```

[[↑] Back to top](#golang-fundamental)

# Reflection

## What is reflection?

Reflection is the ability of a program to inspect its variables and values at run time and find their type. You might not understand what this means but that's alright.

[[↑] Back to top](#golang-fundamental)

# Others

## Parallel Execution

Parallel operation means that two computations are literally running simultaneously - at the same time. At one point in time both computations are advanced. There is no taking turns, they are advanced at the same time.

![48967958-2f20cb00-f023-11e8-80e0-32eff3ac4554](https://user-images.githubusercontent.com/11765228/48969447-95184d00-f039-11e8-876c-5c78adec5cc2.png)

### Why Use Parallel Execution

- Tasks may complete more quickly
- Some tasks are parallelizable and some are not

## Concurrent Execution

Concurrent operation means that two computations can both make progress and advance regardless of the other. If there are two threads, for example, then both make progress independently. The second computation doesn't need to wait for the first one to complete before it can be advanced.

![48967954-23350900-f023-11e8-9f20-ada1ca32ed0c](https://user-images.githubusercontent.com/11765228/48969448-96e21080-f039-11e8-8069-33379bd10ece.png)

## Concurrent vs Parallel

- Parallel tasks must be executed on different hardware
- Concurrent tasks may be executed on the same hardware

## Threads vs. Processes

- Threads share some context
- Many threads can exit in on process
- OS Schedules threads rather than the processes

## Stack vs. Heap

### Stack

Stack is dedicated to function calls.

- Local variables are stored here
- Deallocated after unction complete

### Heap

Heap is persistent

## Garbage Collection

- Go is a compiled language which enables garbage collection
- Implementation is fast
- The compiler determines stack vs heap
- Garbage collection in the background

## Operating System

- CPU
- Threads
- Queues

[[↑] Back to top](#golang-fundamental)
