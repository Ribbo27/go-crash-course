package main

import "fmt"

func main() {

	ids := []int{33,76,54,23,11,2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids togheter
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)

	// Range with map
	emails := map[int]string {
		1: "mike@gmail.com",
		2: "bob@gmail.com",
		3: "sharon@gmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%d: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Key:", k)
	}
}