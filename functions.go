package main

import "fmt"

/*
	Syntax for defining functions :
	func nameOfthefunction (variable datatype,var datatype...) (return type/type's) {
		code

		return statement;(can be omitted)
	}
*/

func plus(a int, b int) int {

	return a + b
}

//When multiple parameters with the same datatype , writing datatype again and again can be avoided

func plusPlus(a, b, c int) int {
	return a + b + c
}

//Go has a unique feature of returning multiple values just need to define the return datatype's in closed brackets

func vals() (int, int) {
	return 3, 7
}

//Variadic functions accept a variable number of arguments/parameters

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//If you only want a subset of the returned values, use the blank identifier _
	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	//If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
