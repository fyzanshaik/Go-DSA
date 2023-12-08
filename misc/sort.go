package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sort.Ints(slice)
	fmt.Println(slice)
}
