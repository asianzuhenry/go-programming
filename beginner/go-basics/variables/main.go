package main

import "fmt"

func main() {

	// datatypes and variable declarations
	// string, int, float64, bool

	// explicitly declaring a variable
	var name string = "Gopher"
	var age int = 5
	var isStudent bool = false
	var gpa float64 = 10.5

	// type inference
	country := "GoLand"
	height := 10.5

	// short variable declaration
	isGopher := true

	// multi-variable declaration
	var x, y, z int = 1, 2, 3

	// printing multiple variables
	fmt.Println("Coordinates:", x, y, z)

	// blank identifier to ignore a variable
	_, b, _ := 4, 5, 6
	fmt.Println("Value of b:", b)

	// multiple variable declaration with type inference
	a, b1, c := 7, 8.5, "Golang"
	fmt.Println("Values:", a, b1, c)

	// constant declaration
	const pi = 3.14

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
	fmt.Println("Height:", height)
	fmt.Println("Is Gopher:", isGopher)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("GPA:", gpa)
	fmt.Println("Pi:", pi)
}
