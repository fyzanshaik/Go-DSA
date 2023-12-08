package main

import "fmt"

func Maximum(x, y int) int {
	if x < y {
		return y
	} else if x > y {
		return x
	}

	return -1
}

func Swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var test, a, b, r int
	fmt.Scan(&test)
	fmt.Scan(&a, &b)

	switch test {
	case 1:
		r = Maximum(a, b)
		fmt.Println(r)
	case 2:
		Swap(&a, &b)
		fmt.Println(a, b)
	default:
		fmt.Println("Invalid test option")
	}
}
