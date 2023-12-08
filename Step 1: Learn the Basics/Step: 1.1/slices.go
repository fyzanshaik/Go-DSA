package main

import (
	"fmt"
)

func main() {
	/*
	Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 
	An uninitialized slice equals to nil and has length 0.
	*/

	// Declare an uninitialized slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Make a slice with a length of 3
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set values in the slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Display the length of the slice
	fmt.Println("len:", len(s))

	// Append elements to the slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copy the slice to a new slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Create slices with different ranges
	//Slices support a “slice” operator with the syntax slice[low:high] it excludes the end index.
	l := s[2:5]
	fmt.Println("sl1:", l)

	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l)
	
	//And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize a slice using shorthand syntax
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Create a 2D slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
