<h1 align="center">Golang Fundamental</h1>

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
