/*
	Reflection
		- The reflect package helps to identify
			the underlying concrete type and the value of a interface{} variable.
*/
package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	data := employee{
		name:    "max",
		id:      1,
		address: "taoyuan",
		salary:  100000,
		country: "taiwan",
	}

	fmt.Printf("reflect.TypeOf(data) => %v\n", reflect.TypeOf(data))
	fmt.Printf("reflect.TypeOf(data).Name() => %v\n", reflect.TypeOf(data).Name())
	fmt.Printf("reflect.ValueOf(data) => %v\n", reflect.ValueOf(data))
	fmt.Printf("reflect.ValueOf(data).Kind() => %v\n", reflect.ValueOf(data).Kind())
	fmt.Printf("reflect.ValueOf(data).NumField() => %v\n", reflect.ValueOf(data).NumField())
	fmt.Printf("reflect.ValueOf(data).Field(0) => %v\n", reflect.ValueOf(data).Field(0))
	fmt.Printf("reflect.ValueOf(data).Field(0).Kind() => %v\n", reflect.ValueOf(data).Field(0).Kind())
}
