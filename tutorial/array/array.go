package main

import "fmt"

func main() {
	basic()
}

func basic() {
	var numbers [5]int
	fmt.Printf("%#v", numbers) // int{0, 0, 0, 0, 0}

	// string literal
	weeks := [5]string{"first", "second", "third", "fourth", "fifth"}
	fmt.Printf("%#v", weeks) // [5]string{"first", "second", "third", "fourth", "fifth"}

	// string literal with ellipsis
	months := [...]string{"September", "October", "November"}
	fmt.Printf("%#v", months) // [3]string{"September", "October", "November"}
}
