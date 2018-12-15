package main

import "fmt"

func main() {
	closure()
}

/*
	CLOSURE
*/
func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func closure() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}
