package main

import "fmt"

func main() {
	// nil slices
	var s1 []int
	fmt.Println(s1)        // []
	fmt.Println(len(s1))   // 0
	fmt.Println(cap(s1))   // 0
	fmt.Println(s1 == nil) // true
	// make slice
	s2 := make([]int, 0)
	fmt.Println(s2)        // []
	fmt.Println(len(s2))   // 0
	fmt.Println(cap(s2))   // 0
	fmt.Println(s2 == nil) // false
	// slice literal
	s3 := []int{}
	fmt.Println(s3)        // []
	fmt.Println(len(s3))   // 0
	fmt.Println(cap(s3))   // 0
	fmt.Println(s3 == nil) // false
}
