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

## Names

In Go language, a name is exported if it starts with capital letter. Exported means the function or variable/constant is accessible to the importer of the respective package.

## Variables

Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in go.

```go
// Declaring a single variable
var name type
// Declaring a variable with initial value
var name type = initialvalue
// Multiple variable declaration
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

A pointer is a variable which stores the memory address of another variable.

![pointer](https://user-images.githubusercontent.com/11765228/48991422-ceea7000-f16d-11e8-955d-d035380500e7.png)

```go
func main() {
	variable := 20
	// pointer to an integer
	var pointer *int
	// & -> get the address of a variable
	pointer = &variable
	// * -> get the value from a pointer variable
	fmt.Println(*pointer)
}

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
var variable_name [SIZE] variable_type
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

![golang-slices-length-capacity](https://user-images.githubusercontent.com/11765228/48991986-9ac47e80-f170-11e8-9a79-cb0509d90c44.jpg)

```go
func main() {
	// a slice of unspecified size
	var numbers []int
	// slice literal
	x := []int{1, 2, 3, 4, 5}
	// len() -> the elements presents in the slice
	fmt.Println(len(numbers))
	// cap() -> the capacity of the slice
	fmt.Println(cap(x))
}
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
func main() {
	// declare a variable, by default map will be nil
	var countryCapitalMap map[string]string
	// define the map as nil map can not be assigned any value
	countryCapitalMap = make(map[string]string)
	// map literal
	countryCapitalMap2 := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	// delete() function is used to delete an entry from a map.
	delete(countryCapitalMap2, "France")
	fmt.Println(countryCapitalMap)
	fmt.Println(countryCapitalMap2)
}
```

## Structs

A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity.

```go
// defining a structs
type person struct {
	name string
	age  int
}

func main() {
	// struct literal
	max := person{
		name: "max",
		age:  24,
	}
	// accessing struct members
	fmt.Println(max.name)
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

## Error

```go
func Sqrt(value float64)(float64, error) {
   if(value < 0) {
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value)
}
```

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

[[↑] Back to top](#golang-notes)

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
