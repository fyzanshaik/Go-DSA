package main

import "fmt"

func modifyValueByReference(x *int) {
	print(x)
	*x = 42
}

func main() {
	value := 10
	fmt.Println("Before: ", value)
	print(&value)
	println()
	modifyValueByReference(&value)
	println()
	fmt.Println("After: ", value)
}
