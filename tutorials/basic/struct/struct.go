/*
	Struct
		- Typed collections of fields
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	basic()
	structLiterals()
	comparingEmptyStruct()
	assignAndAccessStructField()
	embeddingStructs()
}

/*
	THE BASIC
*/
func basic() {
	type myStruct struct {
		name      string
		nickNames []string
		data      map[string]interface{}
	}

	var s1 myStruct
	fmt.Printf("%+v\n", s1)                   // {name: nickNames:[] data:map[]}
	fmt.Printf("%#v\n", s1.nickNames)         // []string(nil)
	fmt.Println(reflect.TypeOf(s1.nickNames)) // []string
	fmt.Printf("%#v\n", s1.data)              // map[string]interface {}(nil) aka nil slice
	fmt.Println(reflect.TypeOf(s1.data))      // map[string]interface {} aks nil map
}

/*
	STRUCT LITERALS
*/
func structLiterals() {
	var person = struct {
		name string
		age  int
	}{
		name: "max",
		age:  24,
	}
	fmt.Printf("%+v\n", person)
}

/*
	Comparing Empty Struct
*/
func comparingEmptyStruct() {
	type myStruct struct{ name string }

	var s1 myStruct
	s2 := myStruct{}
	fmt.Printf("%p\n", &s1)        // 0x40e128
	fmt.Printf("%p\n", &s2)        // 0x40e130
	fmt.Printf("%t\n", &s1 == &s2) // false
	fmt.Printf("%t\n", s1 == s2)   // true

	s1.name = "max"
	fmt.Printf("%p\n", &s1)        // 0x40e128
	fmt.Printf("%p\n", &s2)        // 0x40e130
	fmt.Printf("%t\n", &s1 == &s2) // false
	fmt.Printf("%t\n", s1 == s2)   // false
}

/*
	ACCESS STRUCT FIELDS
	ASSIGN STRUCT FIELDS
*/
func assignAndAccessStructField() {
	type myStruct struct {
		name string
	}

	var person myStruct
	person.name = "max"
	fmt.Println(person.name) // max

	// pointer
	pointer := &person
	fmt.Printf("%p\n", pointer) // 0x442280
	(*pointer).name = "tom"
	fmt.Println((*pointer).name) // tom

	// assign to a struct field through the pointer
	// access struct fields through the pointer
	pointer.name = "nick"
	fmt.Println(pointer.name) // nick
}

/*
	EMBEDDING STRUCTS
*/
type Address struct {
	Street string
	City   string
}
type Employee struct {
	Name   string
	Salary string
	// ANONYMOUS STRUCT FIELDS
	Address
}

func (a *Address) GetFullAddress() string {
	return fmt.Sprintf("I live in %v city and the street is %v", a.City, a.Street)
}

func embeddingStructs() {
	var e1 Employee
	fmt.Printf("%+v\n", e1)
	e1.City = "New York"
	e1.Street = "Broadway"
	fmt.Printf("%+v\n", e1)
	/*
		A type that is storted with a struct type using a anonymous field is said to be ebedded within the struct.
		Methods of an embedded type get promted to the otuer type.
		They can be called as if they were defined on the outer type.
	*/
	fmt.Printf("%+v\n", e1.GetFullAddress())
}
