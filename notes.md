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

[[↑] Back to top](#golang-notes)

## Functions

- Value Argument
- Pointer Argument

[[↑] Back to top](#golang-notes)

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
