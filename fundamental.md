<h1 align="center">Golang Fundamental</h1>

- [Composite Types](#composite-types)
- [Functions](#functions)
- [Methods](#methods)
- [Interfaces](#interfaces)
- [Goroutines and Channels](#goroutines-and-channels)
- [Concurrency with Shared Variables](#concurrency-with-shared-variables)
- [Packages and the Go Tool](#packages-and-the-go-tool)
- [Testing](#testing)
- [Reflection](#reflection)
- [Others](#others)

# Composite Types

## Arrays

The type [n]T is an array of n values of type T.

```go
func main() {
	// declaring an empty array of strings
	var weeks []string
    fmt.Println(weeks)

	// declaring an array with elements
    days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	// interating through arrays
	for _, v := range days {
		fmt.Println(v)
	}
}
```

## Slices

A "window" on an underlying array.

- `Pointer` indicate the start of the slice
- `Length` is the number of elements in the clice
- `Capacity` is the maximum number of elements

![golang-slices-length-capacity](https://user-images.githubusercontent.com/11765228/48991986-9ac47e80-f170-11e8-9a79-cb0509d90c44.jpg)

```go
func main() {
	// a slice of unspecified size
	var numbers []int
	// slice literal
    x := []int{1, 2, 3, 4, 5}
    // making slices
    cities := make([]string, 3)

	// len() -> the elements presents in the slice
	fmt.Println(len(numbers))
	// cap() -> the capacity of the slice
    fmt.Println(cap(x))

    // appending to a slice
    cities = append(cities, "San Diego", "Mountain View")
    fmt.Println(cities)
}
```

## Composition (Struct Embedding) vs inheritance

Coming from an OOP background a lot of us are used to inheritance, something that isn’t supported by Go. Instead you have to think in terms of composition and interfaces.

```go
package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func NewPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		User:   &User{id, name, location},
		GameId: gameId,
	}
}

func main() {
	p := NewPlayer(42, "Matt", "LA", 90404)
	fmt.Println(p.Greetings())
}
```

# Functions

A function is a block of code that performs a specific task. A function takes a input, performs some calculations on the input and generates a output.

```go
func function_name( [parameter list] ) [return_types] {
   // body of the function
}
```

## Call by Value vs. Reference

By default, Go uses call by value to pass arguments. In general, it means the code within a function cannot alter the arguments used to call the function.

### Call by Value

- Passed arguments are copied to parameters.
- Modifying parameter has no effect outside the function.

### Call by Reference

- Programmer can pass a pointer as an argument.
- Called function has direct access to caller viable in memory.

## Variadic Functions

A variadic function is a function that can accept variable number of arguments.

```go
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	res := getMax(1, 2, 3, 4, 5)
	fmt.Println(res)
}
```

### Variadic Slice Argument

```go
func main() {
	vslice := []int{1, 2, 3, 4, 5}
	res := getMax(vslice...)
	fmt.Println(res)
}
```

## Deferred Function Calls

Defer statement is used to execute a function call just before the function where the defer statement is present returns.

```go
func main() {
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!")
}
```

## Panic

`panic` and `recover` can be considered similar to try-catch-finally idiom in other languages except that it is rarely used and when used is more elegant and results in clean code.

### When should panic be used?

One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.

## Recover

recover is a builtin function which is used to regain control of a panicking goroutine.

## Functions are First-class

- Functuns can be treated like other types

### Variables as Functions

```go
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar = incFn
	fmt.Print(funcVar(1))
}
```

### Functions as Arguments

```go
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))
}
```

### Anonymous Functions

```go
func addOne() func() int {
  var x int
  // Anonymous Functions
  return func() int {
    x++
    return x + 1
  }
}

func main() {
  myFunc := addOne()
  fmt.Println(myFunc()) // 2
  fmt.Println(myFunc()) // 3
  fmt.Println(myFunc()) // 4
  fmt.Println(myFunc()) // 5
}
```

### Functions as Return Values

```go
package main

import (
	"fmt"
)

func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	var add2 = add(2)
	fmt.Println(add2(3))
}
```

### Closure

- Function + its environment
- When functions are passed/returned, their environment comes with them

## Guidelines for Functions

### Function Naming

- Give functions a good name
- Behavior can be understood at a glance
- Parameter naming counts too

### Functional Cohesion

- Functions should perform only one "operation"
- An "operation" depends on the context

### Few Parameters

- Debugging requires tracing function input data
- More difficult with a large number of parameters
- Function may have bad functional cohesion

[[↑] Back to top](#golang-fundamental)

# Methods

A method is a function associated with a particular type.

```go
package main

import "fmt"

type myInt int

func (mi myInt) double() int {
	return int(mi * 2)
}

func main() {
    v := myInt(3)
    // call by value
	fmt.Println(v.double())
}
```

## Controlling Access to Structs

```go
package data

// Point struct
type Point struct {
	x float64
	y float64
}

// InitMe initialize value
func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = xn
}
```

```go
package main

import "data"

func main () {
	var p data.Point
	p.initMe(3, 4)
}
```

## Methods with a Pointer Receiver

- The receiver can be a pointer to a type.
- Call by reference, pointer is passed to the method.

### No Need to Reference

```go
p := Point(3, 4)
p.OffsetX(5)
fmt.Println(p.x)
```

- Do not need to reference when calling the method.

### No Need to Dereference

```go
func (p *Point) OffsetX(v int) {
    p.x = p.x + v
}
```

- Point is referenced as p, not \*p.
- Dereferencing is automatic with `.` operator.

[[↑] Back to top](#golang-fundamental)

# Interfaces

An interface type is defined by a set of methods. A value of interface type can hold any value that implements those methods.

```go
type Employee interface {
	Name() string
	Language() string
	Age() int
	Random() (string, error)
}
```

## Interface vs. Concrete Types

### Concrete Types

- Specify the exact representation of the data and methods
- Complete method implementation is included

## Interface Types

- Specifies some method signatures
- Implementations are abstracted

### Interface Values

- Can be treated like other values

### Interface values have two components

1. Dynamic Type: Concrete type which it is assigned to
2. Dynamic Value: value of the dynamic type

## Empty Interface

An interface which has zero methods is called empty interface. It is represented as `interface{}`. Since the empty interface has zero methods, all types implement the empty interface.

```go
func PrintMe(val interface{}) {
    fmt.PrintIn(val)
}
```

## Nil Dynamic Type

```
var s1 Speaker
```

- Cannot call a method, runtime error

## Errors

Many Go programs return error interface objects to indicate errors

```go
type error interface {
    Error() string
}
```

```go
package main

import (
    "fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
```

## Polymorphism

A variable of type interface can hold any value which implements the interface.

```go
/* define an interface */
type Shape interface {
   area() float64
}

/* define a circle */
type Circle struct {
   x,y,radius float64
}

/* define a rectangle */
type Rectangle struct {
   width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
   return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}

func main() {
   circle := Circle{x:0,y:0,radius:5}
   rectangle := Rectangle {width:10, height:5}

   fmt.Printf("Circle area: %f\n",getArea(circle))
   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}
```

## Type Assertions

- Type assertions can be used to detemine and extract the underlying concrete type.

```go
func DrawShape(s Shape2D) bool {
    rect, ok := s.(Rectangle)
    if ok {
        DrawRect(rect)
    }
    tri, ok := s.(Triangle)
    if ok {
        DrawTri(tri)
    }
}
```

## Type Switch

- Switch statement used with a type assertion

```go
func DrawShape (s Shape2D) bool {
    switch sh := s.(type) {
        case Rectangle:
            DrawRact(sh)
        case Triangle:
            DrawTri(sh)
    }
}

```

[[↑] Back to top](#golang-fundamental)

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

# Packages and the Go Tool

## Import Paths

```go
import (
    "fmt"
    "math/rand"
    "encoding/json"

    "golang.org/x/net/html"

    "github.com/go-sql-driver/mysql"
)
```

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
