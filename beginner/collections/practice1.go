// practice on collections in Go
// question: Create a slice of integers, append some values to it,
// and then create a map that associates each integer with its square. F
// inally, print the contents of the map.

package main

import "fmt"

func practice1() {
	numbers := []int{1, 2, 3} // slice of numbers

	// new slice
	numbers = append(numbers, 4, 6, 7, 8)

	squares := make(map[int]int)
	
	for _, num := range numbers {
		squares[num] = num * num
	}

	fmt.Println(squares)

}