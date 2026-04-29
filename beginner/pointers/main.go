package main

import "fmt"

// This program demonstrates the use of pointers in Go.
// It defines a main function that creates a pointer to an integer and modifies its value through the pointer.
// Pointers are variables that store the memory address of another variable.
func main() {
	var x int = 10
	var p *int = &x // p is a pointer to x

	fmt.Println("Value of x:", x)   // Output: Value of x: 10
	fmt.Println("Address of x:", &x) // Output: Address of x: 0xc0000140a8 (example address)
	fmt.Println("Value of p:", p)   // Output: Value of p: 0xc0000140a8 (same as address of x)
	fmt.Println("Value pointed to by p:", *p) // Output: Value pointed to by p: 10

	*p = 20 // Modifying the value of x through the pointer

	fmt.Println("New value of x:", x) // Output: New value of x: 20	
}