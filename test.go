package main

import "fmt"

func main() {
	// declare a variable, by default map will be nil
	var countryCapitalMap map[string]string
	// define the map as nil map can not be assigned any value
	countryCapitalMap = make(map[string]string)
	// map literal
	countryCapitalMap2 := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	// delete() function is used to delete an entry from a map.
	delete(countryCapitalMap2, "France")
	fmt.Println(countryCapitalMap)
	fmt.Println(countryCapitalMap2)
}
