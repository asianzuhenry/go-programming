package main

import "fmt"

func main()  {
	// Create a slice of strings
	// slices are like dynamic arrays, they can grow and shrink in size
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	for i, name := range names {
		// %d is for integers, %s is for strings
		fmt.Printf("Person %d: %s\n", i+1, name)
	}

	names = append(names, "Eve") // Append a new name to the slice
	fmt.Println("After appending Eve:")
	for i, name := range names {
		fmt.Printf("Person %d: %s\n", i+1, name)
	}

	// Create a map to store ages
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