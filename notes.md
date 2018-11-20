<h1 align="center">Golang Notes</h1>

- [Program Structure](#program-tructure)
  - [Variables](#variables)
  - [Type Declarations](#type-declarations)
  - [Scope](#scope)
- [Basic Data Types](#basic-data-types)
- [Composite Types](#composite-types)
  - [Arrays](#arrays)
  - [Slices](#slices)
  - [Maps](#maps)
  - [Structs](#structs)
  - [JSON](#json)
- [Functions](#functions)
- [Pointers](#pointers)
- [Flow Control](#flow-control)
- [Type System](#type-system)
- [Concurrency](#concurrency)
- [Packages](#packages)
- [Testing](#testing)
- [Others](#others)
  - [Operating System](#operating-system)

## Program Structure

### Variables

- `var`
  - Zero Values
- Short Variable Declarations
  - Type Inference

[[↑] Back to top](#golang-notes)

### Type Declarations

> type name underlying-type

[[↑] Back to top](#golang-notes)

### Scope

The scope of a declaration is a region of the program text; it is a compile-time property.

Lexical blocks

Universe block

[[↑] Back to top](#golang-notes)

## Basic Data Types

- Integers
- Floating-Point Numbers
- Complex Numbers
- Booleans
- Strings
- Constants

[[↑] Back to top](#golang-notes)

## Composite Types

### Arrays

```go
[n]T
```

An array is a ﬁxed-length sequence of zero or more elements of a particular type. Because of their ﬁxed length, arrays are rarely used directly in Go. Slices, which can grow and shrink, are much more versatile.

[[↑] Back to top](#golang-notes)

### Slices

```go
[]T
```

Slices represent variable-length sequences whose elements all have the same type.

**Append**

```go
var x []int
x = append(x, 1)
```

[[↑] Back to top](#golang-notes)

### Maps

In Go, a map is a reference to a hash table.

```go
make(map[K]V)
// Map Literal
map[K]V{}
```

[[↑] Back to top](#golang-notes)

### Structs

A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity.

```go
// Struct Literals
type Point struct{ X, Y int }
```

### JSON

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

## Functions

### Function Declarations

```go
func name (parameter-list) (result-list) {
    body
}
```

Arguments are passed by value, so the function receives a copy of each argument; modiﬁcations to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then the caller may be affected by any modiﬁcations the function makes to variables indirectly referred to by the argument.

## Pointers

- `*T`
- `&`
- `*`

[[↑] Back to top](#golang-notes)

## Flow Control

- `defer`
  - Stacking defers

[[↑] Back to top](#golang-notes)

## Type System

- Methods
  - Value Receivers
  - Pointer Receivers
- Interfaces
  - Empty interface
  - Method Sets
  - Polymorphism
- Type Embedding
  - Inner Type Promotion
- Identifiers
  - Exported
  - Unexported

[[↑] Back to top](#golang-notes)

## Concurrency

- Goroutines
  - Select
- Race conditions
  - Race Detector
  - Atomic Functions
  - Mutexes
    - Lock
    - Unlock
- Channels
  - Unbuffered Channels
  - Buffered Channels

[[↑] Back to top](#golang-notes)

## Packages

- `import`

[[↑] Back to top](#golang-notes)

## Testing

[[↑] Back to top](#golang-notes)

## Others

### Operating System

- CPU
- Threads
- Queues

[[↑] Back to top](#golang-notes)
