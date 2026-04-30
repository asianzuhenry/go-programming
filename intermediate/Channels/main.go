package main

// channels
import (
	"fmt"
)

func main() {
	fmt.Println("Hello, GO!")
	var c= make(chan string)
	go func() {
		c <- "Hello from a goroutine!"
	}()
	message := <-c
	fmt.Println(message)
}

// TODO:
// more practice needed to understand goroutines and how they work in Go. 
// Consider exploring channels for communication between goroutines and using sync. 
// WaitGroup to manage the lifecycle of goroutines more effectively.