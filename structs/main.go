package main

import "fmt"

// Person has a firstName & lastName
type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{"Alex", "Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	// if num := 9; num > 0 {
	// 	fmt.Printf("%+v", num)
	// }
	// if num := 9; num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	arr := make([]int, 0, 10)
	fmt.Println(arr)
	name := "Test"

	namePointer := &name
	fmt.Printf("Address of name: %v", namePointer)
	fmt.Printf("  ##  Value at namePointer: %v", *namePointer)

	fmt.Printf("Address of namePointer: %v", &namePointer)
	// fmt.Println(namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
