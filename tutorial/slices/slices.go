package main

import (
	"fmt"
	"reflect"
)

func main() {
	basic()
}

/*
	THE BASIC
*/
func basic() {
	// nil slices
	var s1 []int
	fmt.Println(s1)                 // []
	fmt.Printf("%#v\n", s1)         // []int(nil)
	fmt.Println(reflect.TypeOf(s1)) // []int
	fmt.Println(s1 == nil)          // true
	fmt.Println(len(s1))            // 0
	fmt.Println(cap(s1))            // 0

	// make slice
	s2 := make([]int, 0)
	fmt.Println(s2)                 // []
	fmt.Printf("%#v\n", s2)         // []int{}
	fmt.Println(reflect.TypeOf(s2)) //[]int
	fmt.Println(s2 == nil)          // false
	fmt.Println(len(s2))            // 0
	fmt.Println(cap(s2))            // 0

	// slice literal
	s3 := []int{}
	fmt.Println(s3)                 // []
	fmt.Printf("%#v\n", s3)         // []int{}
	fmt.Println(reflect.TypeOf(s3)) // []int
	fmt.Println(s3 == nil)          // false
	fmt.Println(len(s3))            // 0
	fmt.Println(cap(s3))            // 0
}
