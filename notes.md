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

```go
[n]T
```

An array is a ﬁxed-length sequence of zero or more elements of a particular type. Because of their ﬁxed length, arrays are rarely used directly in Go. Slices, which can grow and shrink, are much more versatile.

## Slices

```go
[]T
```

Slices represent variable-length sequences whose elements all have the same type.

**Append**

```go
var x []int
x = append(x, 1)
```

## Maps

In Go, a map is a reference to a hash table.

```go
make(map[K]V)
// Map Literal
map[K]V{}
```

## Structs

A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity.

```go
// Struct Literals
type Point struct{ X, Y int }
```

## JSON

Converting a Go data structure like movies to JSON is called marshaling.

```go
data, err := json.MarshalIndent(movies, "", "    ")

if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}

fmt.Printf("%s\n", data)
```

A ﬁeld tag is a string of m associated at compile time with the ﬁeld of a struct.

```go
Year int `json:"released"`
Color bool `json:"color,omitempty"`
```

`json.Unmarshal`

Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.

```go
if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
}
fmt.Println(dat)
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

Arguments are passed by value, so the function receives a copy of each argument; modiﬁcations to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then the caller may be affected by any modiﬁcations the function makes to variables indirectly referred to by the argument.

## Blank Identifier

`_` is know as the blank identifier in Go. It can be used in place of any value of any type.

## Variadic Functions

### What is a variadic function?

A variadic function is a function that can accept variable number of arguments.

### Syntax

If the last parameter of a function is denoted by `...T`, then the function can accept any number of arguments of type `T` for the last parameter.

## Deferred Function Calls

### What is Defer?

Defer statement is used to execute a function call just before the function where the defer statement is present returns.

```go
package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
```

## Panic

### What is panic?

`panic` and `recover` can be considered similar to try-catch-finally idiom in other languages except that it is rarely used and when used is more elegant and results in clean code.

### When should panic be used?

One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.

## Recover

recover is a builtin function which is used to regain control of a panicking goroutine.

[[↑] Back to top](#golang-notes)

# Methods

A method is a function associated with a particular type.

## Methods with a Pointer Receiver

Because calling a function makes a copy of each argument value, if a function needs to update a variable, or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable using a pointer. The same goes for methods that need to update the receiver variable: we attach them to the pointer type.

## Composing Types by Struct Embedding

[[↑] Back to top](#golang-notes)

# Interfaces

## What is an interface?

Interface specifies what methods a type should have and the type decides how to implement these methods.

## Interface Values

Conceptually, a value of an interface type, or interface value, has two components, a concrete type and a value of that type. These are called the interface’s dynamic type and dynamic value.

## Empty Interface

An interface which has zero methods is called empty interface. It is represented as `interface{}`. Since the empty interface has zero methods, all types implement the empty interface.

## The error Interface

```go
type error interface {
    Error() string
}
```

## Polymorphism

A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.

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

[[↑] Back to top](#golang-notes)

# Goroutines and Channels

## Goroutines

A go stateme causes the function to be called in a newly created goroutine. The go statement itself completes immediately:

```go
f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

## Channels

A channel is a communication mechanism that lets one goroutine send values to another goroutine.

```go
ch <- x // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch // a receive statement; result is discarded
```

### Unbuffered Channels

```go
ch := make(chan type)
```

### Buffered Channels

```go
ch := make(chan type, capacity)
```

## Looping in Parallel

### WaitGroup

A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing.

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
