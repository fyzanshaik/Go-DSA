package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Integer Types
	var intVar int = 42
	var int32Var int32 = 42
	var int64Var int64 = 42

	// Floating-Point Types
	var float32Var float32 = 3.14
	var float64Var float64 = 3.14

	// Boolean Type
	var boolVar bool = true

	// String Type
	var stringVar string = "Hello, Go!"

	// Character Type (represented as rune)
	var charVar rune = 'A'

	// Complex Types
	var complex64Var complex64 = 3 + 2i
	var complex128Var complex128 = 3 + 2i

	fmt.Println("Integer Types:")
	fmt.Println("intVar:", intVar, reflect.TypeOf(intVar))
	fmt.Println("int32Var:", int32Var, reflect.TypeOf(int32Var))
	fmt.Println("int64Var:", int64Var, reflect.TypeOf(int64Var))

	fmt.Println("\nFloating-Point Types:")
	fmt.Println("float32Var:", float32Var, reflect.TypeOf(float32Var))
	fmt.Println("float64Var:", float64Var, reflect.TypeOf(float64Var))

	fmt.Println("\nBoolean Type:")
	fmt.Println("boolVar:", boolVar, reflect.TypeOf(boolVar))

	fmt.Println("\nString Type:")
	fmt.Println("stringVar:", stringVar, reflect.TypeOf(stringVar))

	fmt.Println("\nCharacter Type:")
	fmt.Println("charVar:", charVar, reflect.TypeOf(charVar))

	fmt.Println("\nComplex Types:")
	fmt.Println("complex64Var:", complex64Var, reflect.TypeOf(complex64Var))
	fmt.Println("complex128Var:", complex128Var, reflect.TypeOf(complex128Var))
}
