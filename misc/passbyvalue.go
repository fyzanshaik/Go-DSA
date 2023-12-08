package main

import "fmt"

func modifyValue(value int) {
	value = 42
}

func main() {
	value := 10
	fmt.Println("Before: ", value)

	modifyValue(value)

	fmt.Println("After: ", value) // Output: After: 10
}
