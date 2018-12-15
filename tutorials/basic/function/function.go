package main

import "fmt"

func main() {
	variadicFunctions()
	anonymousFunctions()
}

/*
	VARIADIC FUNCTIONS
*/
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func variadicFunctions() {
	res := getMax(1, 2, 3, 4, 5)
	fmt.Printf("%d\n", res)
}

/*
	ANONYMOUS FUNCTIONS
*/
func anonymousFunctions() {
	func() {
		fmt.Println("hello world first class function")
	}()
}
