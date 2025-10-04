package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define command-line flags
    a, b := flag.Int("a", 0, "the first number"), flag.Int("b", 0, "the second number")
    operator := flag.String("operator", "add", "the operator to use (e.g., add, subtract)")

    // Parse the flags
    flag.Parse()

    // Generate message
    var result int
    switch *operator {
    case "add":
        result = *a + *b
    case "subtract":
        result = *a - *b
	case "multiply":
		result = *a * *b
	case "divide":
		if *b == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		result = *a / *b
    default:
        fmt.Println("Unknown operator")
        os.Exit(1)
    }

    // Print the message
    fmt.Println(result)

    // Optional: Exit with a custom status code
    os.Exit(0)
}

// go run cli-calc.go --a=4 --b=7 --operator=multiply
// go run cli-calc.go --a=4 --b=7 --operator=add
// go run cli-calc.go --a=4 --b=7 --operator=subtract
// go run cli-calc.go --a=4 --b=0 --operator=divide
// go run cli-calc.go --a=4 --b=7 --operator=divide
