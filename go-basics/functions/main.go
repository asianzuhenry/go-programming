package main
import "fmt"

func main() {
	Greet("Gopher")
	sum := Add(4, 6)

	fmt.Println("Sum:", sum)
}


func Greet(name string) {
	fmt.Println("Hello,", name)
}


func Add(a int, b int) int {
	return a + b
}