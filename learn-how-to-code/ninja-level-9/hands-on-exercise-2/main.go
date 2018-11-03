package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hey")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		"Huseyin",
	}

	saySomething(&p)
}
