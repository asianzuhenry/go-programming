package main

import "fmt"

func main() {
	var age int = 12
	var hasid = false
	var numbers = []int{10, -5, 0, 3, 8, 95, 22, 15, 5, 35}

	validateVoter(age, hasid)

	for _, num := range numbers {
		checkNumber(num)
		checkEvenOdd(num)
	}

	for _, number := range numbers {
		fmt.Println("Checking number:", number)
	}
}

// validateVoter checks if a person is eligible to vote based on age and ID possession
func validateVoter(age int, hasid bool) {

	if age < 18 && hasid {
		fmt.Println("you are too young to vote")
	} else if age < 18 && !hasid {
		fmt.Println("you are too young to vote, you need an ID")
	} else if age >= 18 && !hasid {
		fmt.Println("you are not a valid voter, you need an ID")
	} else {
		fmt.Println("you are a valid voter")
	}

}


func login(username, password string) {
	if username == "admin" && password == "password" {
		fmt.Println("login successful")
	} else if username == "admin" && password != "password" {
		fmt.Println("invalid password")
	} else if username != "admin" && password == "password" {
		fmt.Println("invalid username")
	} else {
		fmt.Println("invalid username and password")
	}
}

func checkNumber(num int) {
	if num > 0 {
		fmt.Println("the number is positive")
	} else if num < 0 {
		fmt.Println("the number is negative")
	} else {
		fmt.Println("the number is zero")
	}
}

func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

func checkGrade(score int) {
	if score >= 90 {
		fmt.Println("grade A")
	} else if score >= 80 {
		fmt.Println("grade B")
	} else if score >= 70 {
		fmt.Println("grade C")
	} else if score >= 60 {
		fmt.Println("grade D")
	} else {
		fmt.Println("grade F")
	}
}

func checkTemperature(temp int) {
	if temp > 30 {
		fmt.Println("it's hot")
	} else if temp >= 20 {
		fmt.Println("it's warm")
	} else if temp >= 10 {
		fmt.Println("it's cool")
	} else {
		fmt.Println("it's cold")
	}
}

func checkDay(day int) {
	if day == 1 {
		fmt.Println("Monday")
	} else if day == 2 {
		fmt.Println("Tuesday")
	} else if day == 3 {
		fmt.Println("Wednesday")
	} else if day == 4 {
		fmt.Println("Thursday")
	} else if day == 5 {
		fmt.Println("Friday")
	} else if day == 6 {
		fmt.Println("Saturday")
	} else if day == 7 {
		fmt.Println("Sunday")
	} else {
		fmt.Println("invalid day")
	}
}

func checkMonth(month int) {
	if month == 1 {
		fmt.Println("January")
	} else if month == 2 {
		fmt.Println("February")
	} else if month == 3 {
		fmt.Println("March")
	} else if month == 4 {
		fmt.Println("April")
	} else if month == 5 {
		fmt.Println("May")
	} else if month == 6 {
		fmt.Println("June")
	} else if month == 7 {
		fmt.Println("July")
	} else if month == 8 {
		fmt.Println("August")
	} else if month == 9 {
		fmt.Println("September")
	} else if month == 10 {
		fmt.Println("October")
	} else if month == 11 {
		fmt.Println("November")
	} else if month == 12 {
		fmt.Println("December")
	} else {
		fmt.Println("invalid month")
	}
}