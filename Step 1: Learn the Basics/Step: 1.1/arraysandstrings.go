package main

import (
	"fmt"
)

func main() {
	// Arrays: Fixed-size collections of elements of the same type
	var intArray [5]int              // Declaration of an integer array with size 5
	intArray = [5]int{1, 2, 3, 4, 5} // Initialization
	fmt.Println("Array:", intArray)

	// Strings: Immutable sequence of characters
	var greeting string = "Hello, Go!" // Declaration and initialization of a string
	fmt.Println("String:", greeting)

	// Slices: Dynamic and flexible view into arrays
	intSlice := []int{1, 2, 3, 4, 5} // Declaration and initialization of a slice
	fmt.Println("Slice:", intSlice)

	// Slicing a slice (creating a new slice from the original)
	slicedSlice := intSlice[1:4] // [2, 3, 4]
	fmt.Println("Sliced Slice:", slicedSlice)

	// Modifying a slice also modifies the underlying array
	slicedSlice[1] = 999
	fmt.Println("Modified Slice:", intSlice) // [1, 2, 999, 4, 5]

	// Common built-in functions for slices
	fmt.Println("Length of Slice:", len(intSlice))   // Length of the slice
	fmt.Println("Capacity of Slice:", cap(intSlice)) // Capacity of the slice

	// Append function for slices
	intSlice = append(intSlice, 6, 7, 8) // Append elements to the slice
	fmt.Println("Appended Slice:", intSlice)

	// Copy function for slices
	copyOfSlice := make([]int, len(intSlice))
	copy(copyOfSlice, intSlice)
	fmt.Println("Copied Slice:", copyOfSlice)
}
