package main

import "fmt"

func maicxn()  {
	// Collections in Go include slices and maps, which are used to store and manage groups of related data.
	// Create a slice of strings
	// slices are like dynamic arrays, they can grow and shrink in size
	names := []string{"Alice", "Bob", "Charlie", "Diana"} // This is a slice literal, which creates a slice with the specified elements
	for i, name := range names {
		// %d is for integers, %s is for strings
		fmt.Printf("Person %d: %s\n", i+1, name)
	}

	// Slices can be modified
	names = append(names, "Eve") // Append a new name to the slice
	fmt.Println("After appending Eve:")
	for i, name := range names {
		fmt.Printf("Person %d: %s\n", i+1, name)
	}

	// Create a map to store ages
	// Maps are collections of key-value pairs, where each key is unique and maps to a value
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	ages["Diana"] = 28 // Add a new entry to the map

	fmt.Println("Ages:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	
	// Check if a key exists in the map
	// The syntax "value, exists := map[key]" allows us to check if a key exists in the map. 
	// If it does, "exists" will be true and "value" will contain the corresponding value. 
	// If it doesn't, "exists" will be false and "value" will be the zero value for the type (in this case, 0 for int).
	nameToCheck := "Eve"
	if age, exists := ages[nameToCheck]; exists {
		fmt.Printf("%s is %d years old\n", nameToCheck, age)
	} else {
		fmt.Printf("%s is not in the map\n", nameToCheck)
	}
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}