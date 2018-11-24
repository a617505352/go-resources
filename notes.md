<h1 align="center">Golang Notes</h1>

- [Program Structure](#program-tructure)
- [Basic Data Types](#basic-data-types)
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

# Program Structure

## Variables

Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in go.

### Declaring a single variable

```go
var name type
```

### Declaring a variable with initial value

```go
var name type = initialvalue
```

### Multiple variable declaration

```go
var name1, name2 type = initialvalue1, initialvalue2
```

### Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

- `0` for numeric types
- `false` for the boolean type
- `""` (the empty string) for strings

### Short Variable Declarations

```go
name := expression
```

## Pointers

### What is a pointer

A pointer is a variable which stores the memory address of another variable.

### Declaring pointers

`*T` is the type of the pointer variable which points to a value of type `T`.

### Get get the address of a variable

The `&` operator is used to get the address of a variable.

### Dereferencing a pointer

Dereferencing a pointer means accessing the value of the variable which the pointer points to. `*a` is the syntax to deference a.

```go
var x int = 1
var y int
var ip *int // ip is pointer to int

ip = &x     // ip now points to x
y = *ip     // y is now 1
```

## Type Declarations

```go
type name underlying-type
```

## Scope

The places in code where a variable can be accessed.

### Blocks

A sequence of declarations and statements within matching brackets, `{}`

#### Implicit Blocks

- Universe bloack - all Go source
- Packageblock - all source in a package
- File block - all source in a file
- `if`, `for`, `switch` ...

### Lexical Scoping

Lexical Scoping defines how variable names are resolved in nested functions. Other names of Lexical Scoping are Static Scoping or Closure. It means that the scope of an inner function contains the scope of a parent function.

[[↑] Back to top](#golang-notes)

# Basic Data Types

- Integers
- Floating-Point Numbers
- Complex Numbers
- Booleans
- Strings

## Constants

### The Constant Generator _iota_

```go
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

## Type Conversions

`T(v)` is the syntax to convert a value v to type T

[[↑] Back to top](#golang-notes)

# Composite Types

## Arrays

An array is a **ﬁxed-length** sequence of zero or more elements of a particular type.

```go
// [n]T
var x [5]int
```

### Array Literal

```go
// [n]T{element1, element2}
var x [5]int{1, 2, 3, 4, 5}
```

### Interating Through Arrays

```go
x := [3]int {1, 2, 3}

for i, v range x {
    fmt.Printf("index %d, value %d", i, v)
}
```

## Slices

A "window" on an underlying array.

- `Pointer` indicate the start of the slice
- `Length` is the number of elements in the clice
- `Capacity` is the maximum number of elements

```go
// []T
var x []int
```

### Slice Literal

```go
var x []int{1, 2, 3, 4, 5}
```

### Make

The make built-in function allocates and initializes an object of type slice, map, or chan (only).

```go
person := make(map[string]string)
receiver := make(chan string)
```

### Append

```go
var x []int
x = append(x, 1)
```

## Maps

In Go, a map is a reference to a hash table.

```go
idMap := make(map[K]V)
```

### Map Literal

```go
map[K]V{}
```

## Structs

A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity.

### Struct Literal

```go
type struct Person {
    firstName string
    lastName
}

max := Person{
    firstName: "max",
    lastName: "Li"
}
```

## JSON

Converting a Go data structure like movies to JSON is called marshaling.

### JSON Marshalling

```go
type person struct {
    Name  string
    Addr  string
    Phone string
}

p1 := person{
    Name:  "joe",
    Addr:  "a st.",
    Phone: "123",
}

barr, err := json.Marshal(p1)
if err != nil {
    panic(err)
}

fmt.Print(string(barr))
```

### JSON Unmarshalling

```go
type person struct {
    Name  string
    Addr  string
    Phone string
}

byt := []byte(`
    {
        "Name":"joe",
        "Addr":"a st.",
        "Phone":"123"
    }
`)
data := person{}

err := json.Unmarshal(byt, &data)
if err != nil {
    panic(err)
}

fmt.Println(data)
```

### ﬁeld tag

A ﬁeld tag is a string of m associated at compile time with the ﬁeld of a struct.

```go
Year int `json:"released"`
Color bool `json:"color,omitempty"`
```

[[↑] Back to top](#golang-notes)

# Functions

## What is a function?

A function is a block of code that performs a specific task. A function takes a input, performs some calculations on the input and generates a output.

## Function Declarations

```go
func functionname(parametername type) returntype {
 // function body
}
```

## Call by Value vs. Reference

### Call by Value

- Passed arguments are copied to parameters.
- Modifying parameter has no effect outside the function.

### Call by Reference

- Programmer can pass a pointer as an argument.
- Called function has direct access to caller viable in memory.

### Function Complexity

- Function length is the most obvious measure

## Blank Identifier

`_` is know as the blank identifier in Go. It can be used in place of any value of any type.

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

### What is Defer?

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

### What is panic?

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
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func main() {
	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(v)
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

[[↑] Back to top](#golang-notes)

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

## Composing Types by Struct Embedding

[[↑] Back to top](#golang-notes)

# Interfaces

Interface specifies what methods a type should have and the type decides how to implement these methods.

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

## The error Interface

Many Go programs return error interface objects to indicate errors

```go
type error interface {
    Error() string
}
```

## Polymorphism

A variable of type interface can hold any value which implements the interface.

```go
package main

import (
	"fmt"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func calculateNetIncome(ic []Income) {
	var netincome int
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
```

The above program will output,

```
Income From Project 1 = $5000
Income From Project 2 = $10000
Income From Project 3 = $4000
Income From Banner Ad = $1000
Income From Popup Ad = $3750
Net income of organisation = $23750
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

[[↑] Back to top](#golang-notes)

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

## Multiplexing with select

### What is select?

The select statement is used to choose from multiple send/receive channel operations.

```go
package main

import (
    "fmt"
    "time"
)

func server1(ch chan string) {
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}
func main() {
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}
```

### Default case

The default case in a select statement is executed when none of the other case is ready. This is generally used to prevent the select statement from blocking.

```go
package main

import (
    "fmt"
    "time"
)

func process(ch chan string) {
    time.Sleep(10500 * time.Millisecond)
    ch <- "process successful"
}

func main() {
    ch := make(chan string)
    go process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default:
            fmt.Println("no value received")
        }
    }

}
```

[[↑] Back to top](#golang-notes)

# Concurrency with Shared Variables

## Mutual Exclusion: sync.Mutex

A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.

```go
mutex.Lock()
x = x + 1
mutex.Unlock()
```

[[↑] Back to top](#golang-notes)

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

## The Package Declaration

```go
package main
```

## Import Declarations

```go
import (
    "crypto/rand"
    mrand "math/rand" // alternative name mrand avoids conflict
)
```

## Blank Imports

```go
import _ "image/png" // register PNG decoder
```

## The Go Tool

The go tool combines the features of a diverse set of tools into one command set.

```bash
go
```

[[↑] Back to top](#golang-notes)

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

[[↑] Back to top](#golang-notes)

# Reflection

## What is reflection?

Reflection is the ability of a program to inspect its variables and values at run time and find their type. You might not understand what this means but that's alright.

[[↑] Back to top](#golang-notes)

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

## Control Flows

### For loop

```go
for initialization; condition; update {
    // zero or more statements
}
```

```go
// a traditional "while" loop
for condition {
    // ...
}
```

```go
// inﬁnite loop
for {
    // ...
}
```

### Switch/Case

```go
switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
}
```

```go
// tagless switch
func Signum(x int) int {
    switch {
        case x > 0:
            return +1
        default: return 0
            case x < 0:
        return -1
    }
}
```

## Comments

- Single-line comments

```go
// This is a comment
var x int // Another comment
```

- Block comments

```go
/* Comment 1 Comment 2 */
var x int
```

## Operating System

- CPU
- Threads
- Queues

[[↑] Back to top](#golang-notes)
