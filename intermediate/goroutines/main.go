package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, GO!")
	go func() {
		fmt.Println("Hello from a goroutine!")
	}()
	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}

// TODO:
// more practice needed to understand goroutines and how they work in Go. 
// Consider exploring channels for communication between goroutines and using sync. 
// WaitGroup to manage the lifecycle of goroutines more effectively.