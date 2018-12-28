/*
	Map
		- A map is a builtin type in Go which associates a value to a key.
		- The value can be retrieved using the corresponding key.
*/
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
	// nil map
	var m1 map[string]interface{}
	fmt.Println(m1)                 // {}
	fmt.Printf("%#v\n", m1)         // map[string]interface {}(nil)
	fmt.Println(reflect.TypeOf(m1)) // map[string]interface {}
	fmt.Println(m1 == nil)          // true

	// make map
	m2 := make(map[string]interface{})
	fmt.Println(m2)                 // {}
	fmt.Printf("%#v\n", m2)         // map[string]interface {}{}
	fmt.Println(reflect.TypeOf(m2)) // map[string]interface {}
	fmt.Println(m2 == nil)          // false

	// map literal
	m3 := map[string]interface{}{}
	fmt.Println(m3)                 // {}
	fmt.Printf("%#v\n", m3)         // map[string]interface {}{}
	fmt.Println(reflect.TypeOf(m3)) // map[string]interface {}
	fmt.Println(m3 == nil)          // false
}
