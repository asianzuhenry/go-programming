package main

import "fmt"

func main() {
    fmt.Println("Hello Gopher!")
    fmt.Println(add(1, 2))           // Output: 3
    fmt.Println(add(1.5, 2.5))       // Output: 4.0
}

func add[T int | float64](a, b T) T {
    return a + b
}

// TODO:
// 1. Implement a generic function to find the maximum of two values.
// 2. Create a generic struct to hold a pair of values and implement a method to swap them.
// 3. Practice using generics with slices by creating a function that takes a slice of any type and returns the first element.