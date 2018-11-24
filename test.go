package main

import (
	"fmt"
)

type test struct {
}

func (t *test) what(value int) {
	fmt.Print("11")
}

func main() {
	var wg test
	wg.what(1)
}
