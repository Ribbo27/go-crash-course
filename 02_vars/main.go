package main

import "fmt"

var var1 = "Global var"

func main() {

	// Using var
	var age int32 = 37
	const isCool = true
	
	// Shorthand
	//name := "Ribbo"
	//size := 1.3

	name, size := "Ribbo", 1.3

	fmt.Println(name, age, isCool, var1, size)
	fmt.Printf("%T\n", size)
}