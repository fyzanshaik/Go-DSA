package main

import "fmt"

func main() {
	/*
		Variables can be declared by using var or const keywords
		Go will automatically infer initialized values
		For dynamic variables we need to specify their data type by adding it after their name
		:= is a shorthand for declaring and initializing variables
	*/


    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}