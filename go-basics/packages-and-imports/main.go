package main


import (
	"fmt"
	"math"
	"strings"
)

// ExampleFunction demonstrates usage of imported packages.
func ExampleFunction() {
	// Using the math package to calculate the square root
	number := 16.0
	sqrt := math.Sqrt(number)
	// Using the fmt package to print the result
	fmt.Printf("The square root of %.2f is %.2f\n", number, sqrt) 

	// Using the strings package to manipulate strings
	text := "hello, world"
	upperText := strings.ToUpper(text)
	fmt.Printf("Uppercase: %s\n", upperText)
}

func main() {
	ExampleFunction()
}

// fmt and math are standard library packages in Go.
// 1. fmt is used for formatted I/O operations like printing to the console.
// 2. math provides basic constants and mathematical functions. eg. Sqrt for square root, Pow for power calculations. and so on.
// 3. strings is another standard library package used for string manipulation. It provides functions like ToUpper, ToLower, Contains, and many more.