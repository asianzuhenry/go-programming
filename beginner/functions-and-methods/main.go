package main
import "fmt"

// This program demonstrates the use of functions and methods in Go.
// It defines a main function that calls two other functions: Greet and Add.
func main() {
	Greet("Gopher")
	sum := Add(4, 6)

	fmt.Println("Sum:", sum)
}
// methods are functions that are associated with a specific type. 
// In Go, you can define methods on any type, including structs. 
// However, in this example, we are only using regular functions for simplicity.

func Greet(name string) {
	fmt.Println("Hello,", name)
	p.Greet() // This will call the Greet method on the Person struct, printing "Hello, my name is Alice"
}


func Add(a int, b int) int {
	return a + b
}

//methods are defined with a receiver, which is the type that the method is associated with. 
// For example, if we had a struct type called "Person", we could define a method on it like this:

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

var p Person = Person{Name: "Alice", Age: 30}