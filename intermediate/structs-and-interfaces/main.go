package main

import "fmt"

// struct
type User struct {
	username string
	email string
	number int
	password string
}

// interface
func greetUser(us User) string {
	return "Hi " + us.username
}

func main() {
	fmt.Println("Hello, GO!")

	var U User = User{username: "Henry", email: "henry@gmail.com", number: 01234433, password: "123456789"}
	fmt.Println(U.username)

	fmt.Println(greetUser(U))

}

// TODO: 
// more work needed on structs and interfaces