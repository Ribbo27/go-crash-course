package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Bob",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "m",
		age:       25,
	}

	// Alternative
	person2 := Person{"Samantha", "Williams", "Boston", "f", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// fmt.Println(person2)
	// fmt.Println(person2.age)
	// person2.age++
	// fmt.Println(person2)
	// fmt.Println(person1.greet())
	person1.hasBirthday()
	person2.getMarried(person1.lastName)
	fmt.Println(person2.greet())
}
