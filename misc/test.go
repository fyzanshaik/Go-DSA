package main

import "fmt"

func main() {
	// pattern(6)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum(nums...)

}
	
func pattern(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i+1; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func sum(nums ...int) {
	total := 0
	for _, element := range nums {
		total += element
	}
	fmt.Println(total)
}
