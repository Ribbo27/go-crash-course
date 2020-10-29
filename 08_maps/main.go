package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[int]string)

	// Assign key values
	// emails[1] = "bob@gmail.com"
	// emails[2] = "sharon@gmail.com"
	// emails[3] = "mike@gmail.com"

	emails := map[int]string {
		1: "mike@gmail.com",
		2: "bob@gmail.com",
		3: "sharon@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[1])
	
	// Delete from map
	delete(emails, 2)
	
	fmt.Println(emails)
}
