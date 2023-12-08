package main

import "fmt"

func main() {
    // Read from console
    var input string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&input)
    fmt.Printf("Hello, %s!\n", input)

    // Read multiple values
    var num1, num2 int
    fmt.Print("Enter two numbers separated by space: ")
    fmt.Scanf("%d %d", &num1, &num2)
    sum := num1 + num2
    fmt.Printf("Sum: %d\n", sum)
}
