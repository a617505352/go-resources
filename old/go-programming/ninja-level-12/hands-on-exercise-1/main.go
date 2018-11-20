package main

import (
	"fmt"

	"github.com/LIYINGZHEN/learning-go/learn-how-to-code/ninja-level-12/hands-on-exercise-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
