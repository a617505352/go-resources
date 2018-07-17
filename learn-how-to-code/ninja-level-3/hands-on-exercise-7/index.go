package main

import "fmt"

func main() {
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDBONDBONDBONDBOND", x)
	} else {
		fmt.Println("neither")
	}
}
