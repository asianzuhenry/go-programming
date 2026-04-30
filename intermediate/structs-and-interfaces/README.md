# Structs and Interfaces

In this section, we will explore the concepts of structs and interfaces in Go. Structs are used to create complex data types that group together related data, while interfaces allow us to define behavior that can be implemented by different types.

## Structs
A struct is a composite data type that groups together variables under a single name. Each variable in a struct is called a field. Here's an example of how to define and use a struct in Go:

```go
package main
import "fmt"
type Person struct {
    Name string
    Age  int
}
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name) // Output: Alice
    fmt.Println(p.Age)  // Output: 30
}
```

In this example, we defined a struct called `Person` with two fields: `Name` and `Age`. We then created an instance of `Person` and accessed its fields.

## Interfaces
An interface in Go is a type that defines a set of method signatures. A type that implementss all the methods of an interface is said to satisfy that interface. Here's an example of how to define and use an interface in Go:
```go 
package main
import "fmt"
type Shape interface {
    Area() float64
}
type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
func main() {
    c := Circle{Radius: 5}
    var s Shape = c
    fmt.Println(s.Area()) // Output: 78.5
}
```

In this example, we defined an interface called `Shape` with a method `Area()`. We then defined a struct `Circle` that implements the `Area()` method. Finally, we created an instance of `Circle`, assigned it to a variable of type `Shape`, and called the `Area()` method.

In summary, structs and interfaces are fundamental concepts in Go that allow us to create complex data types and define behavior. Structs help us group related data together, while interfaces enable us to define contracts for behavior that can be implemented by different types.

