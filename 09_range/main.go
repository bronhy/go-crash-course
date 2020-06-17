package main

import "fmt"

func main() {
	// slice of numbers
	ids := []int{33, 44, 55, 66, 77}

	// Loop through ids index, value
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Use _ for not used variables
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// get only the key
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
