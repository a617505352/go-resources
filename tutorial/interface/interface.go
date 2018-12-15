/*
- An interface is a set of methods that certain values are expected to have.
- Any type that has all the methods listed in an interface definition
	is said to satisfy that interface.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	pointerMethods()
	typeAssertions()
	typeSwitch()
	polymorphism()
	emptyInterface()
	errorInterface()
}

/*
	POINTER METHONDS

	If a type declares methods with pointer receivers,
		then you'll only be able to use pointers to that type when assigning to interface variables.
*/
type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface{ toggle() }

func pointerMethods() {
	s := Switch("off")
	// var t Toggleable = s
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}

/*
	 Type Assertions
		- When you have a value of a concrete type assigned to a variable with an interface type, 
			a type assertion lets you get the concrete type back.
*/
func typeAssertions() {
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
}

/*
	 Type Switch
		- Switch statement used with a type assertion
*/
func () {
	func DrawShape (s Shape2D) bool {
    switch sh := s.(type) {
        case Rectangle:
            DrawRact(sh)
        case Triangle:
            DrawTri(sh)
    }
}
}

/*
	Polymorphism
		- A variable of type interface can hold any value which implements the interface.
*/
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func polymorphism() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}

/*
	Empty Interface
		- The empty interface, and it's used to accept values of any type
*/
func acceptAnything(thing interface{}) {
	fmt.Printf("%+v", thing)
}

func emptyInterface(){
	acceptAnything("string")
	acceptAnything(1)
}

/*
	Error interface
		- 
*/

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