/*
	Method
		- A method is just a function with a special receiver type that is written between the func keyword and the method name.
*/
package main

import "fmt"

func main() {
	pointerReceiver()
}

/*
	POINTER RECEIVER
*/
type myType string

func (m myType) method() {
	fmt.Printf("%v\n", m)
}
func (m *myType) pointerMethod() {
	fmt.Printf("%v\n", m)  // Hello
	fmt.Printf("%v\n", *m) // 0x40e128
}

func pointerReceiver() {
	value := myType("Hello")
	value.method()
	value.pointerMethod()

	pointer := &value
	pointer.method()
	pointer.pointerMethod()
}
