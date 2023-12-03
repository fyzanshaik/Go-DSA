package main

import "fmt"

func main() {
	/*
		Go does not have while or do while loops , only a for and for each
	*/
	//The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Printf("%v", i)
		i++
	}

	//Classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Print(i, " ")
	}

	//Infinite loop
	for {
		fmt.Print("Infinity\n")
		break
	}

	//For-Range loop used for iterating over elements of various data structures such as arrays slice etc it returns index and value of each element

	sentence := "Shaik"
	for index, char := range sentence {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
	evenNumbers()
}

func evenNumbers() {
	var naturalNumbers [100]int
	var evenNumbers [50]int
	for i := 1; i <= 100; i++ {
		naturalNumbers[i-1] = i
	}

	// Extract and print even numbers from the array
	//In go you can't just not use a variable otherwise it will show error so in order to ignore we can use "_" instead
	evenIndex := 0
	for _, num := range naturalNumbers {
		if num%2 == 0 {
			evenNumbers[evenIndex] = num
			evenIndex++
		}
	}
	fmt.Printf("%v \n %v \n", naturalNumbers, evenNumbers)
}
