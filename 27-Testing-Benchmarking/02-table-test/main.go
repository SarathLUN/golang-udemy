package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", mySum(1, 2))
	fmt.Println("1 + 2 + 3 =", mySum(1, 2, 3))
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
