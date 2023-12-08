package main

import "fmt"

func main() {
	// Output
	fmt.Print("Enter your name: ")

	// Input
	var inputName string
	fmt.Scanln(&inputName)

	// Output using the input
	fmt.Printf("Hello, %s!\n", inputName)

	// Output
	fmt.Print("Enter your age: ")

	// Input
	var inputAge int
	fmt.Scanln(&inputAge)

	// Output using the input
	fmt.Printf("In 10 years, you will be %d years old!\n", inputAge+10)
}
