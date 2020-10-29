package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// asssign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	colorArr := [2]string{"Red","Yellow"}

	fmt.Println(fruitArr)
	fmt.Println(colorArr)
	
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	
	fmt.Println(fruitSlice)
}