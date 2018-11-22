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
  - [Operating System](#operating-system)

# Program Structure

## Variables

```go
var name type = expression
```

Variables declared without a corresponding initialization are zero-valued.

### Short Variable Declarations

```go
name := expression
```

## Pointers

- `*T`
- `&`
- `*`

## Type Declarations

> type name underlying-type

## Scope

The scope of a declaration is a region of the program text; it is a compile-time property.

Lexical blocks

Universe block

[[↑] Back to top](#golang-notes)

# Basic Data Types

- Integers
- Floating-Point Numbers
- Complex Numbers
- Booleans
- Strings
- Constants

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

## Function Declarations

```go
func name (parameter-list) (result-list) {
    body
}
```

Arguments are passed by value, so the function receives a copy of each argument; modiﬁcations to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then the caller may be affected by any modiﬁcations the function makes to variables indirectly referred to by the argument.

# Methods

A method is a function associated with a particular type.

## Methods with a Pointer Receiver

Because calling a function makes a copy of each argument value, if a function needs to update a variable, or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable using a pointer. The same goes for methods that need to update the receiver variable: we attach them to the pointer type.

## Composing Types by Struct Embedding

[[↑] Back to top](#golang-notes)

# Interfaces

Interface types express generalizations or abstractions about the behaviors of other types.

## Interfaces as Contracts

An interface is an abstract type. It doesn’t expose the representation or internal structure of its values, or the set of basic operations they support; it reveals only some of their methods. When you have a value of an interface type, you know nothing about what it is; you know only what it can do, or more precisely, what behaviors are provided by its methods.

## Interface Types

An interface type speciﬁes a set of methods that a concrete type must possess to be considered an instance of that interface.

## Interface Values

Conceptually, a value of an interface type, or interface value, has two components, a concrete type and a value of that type. These are called the interface’s dynamic type and dynamic value.

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

[[↑] Back to top](#golang-notes)

# Concurrency with Shared Variables

## Mutual Exclusion: sync.Mutex

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

[[↑] Back to top](#golang-notes)

# Testing

## Test Functions

```go
func TestName(t *testing.T) {
    // ...
}
```

## Benchmark Functions

```go
func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}
```

[[↑] Back to top](#golang-notes)

# Reflection

[[↑] Back to top](#golang-notes)

# Others

## Operating System

- CPU
- Threads
- Queues

[[↑] Back to top](#golang-notes)
