package main

import "fmt"

func main() {
	practice2()
}

func practice2() {
	// Implementation for practice question 2
	strArr := [3]string{"Go", "is", "fun"}
	fmt.Println(strArr)
	fmt.Println(len(strArr))

	numSlice := []int{1, 2, 3, 4, 5}

	numSlice = append(numSlice, 6, 7)
	fmt.Println(numSlice)
	fmt.Println(len(numSlice))
	fmt.Println(cap(numSlice))

	// Map usage
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	fmt.Println(ages)

	ages["Diana"] = 28
	fmt.Println(ages)

	delete(ages, "Alice")
	fmt.Println(ages)

	arr := [6]int{10, 20, 30, 40, 50, 60}
	fmt.Println(arr)

	slice := arr[1:4] // This creates a slice that includes elements from index 1 to 3 (4 is exclusive)
	fmt.Println(slice)

	var name string = "Bob"
	if ages[name] != 0 {
		fmt.Printf("%s is in the map\n", name)
	} else {
		fmt.Printf("%s is not in the map\n", name)
	}
}