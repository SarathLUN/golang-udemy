package main

import "fmt"

func sumI(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sumI(ii...)
	fmt.Println("sum of all numbers =", s)

	s2 := sumEven(sumI, ii...)
	fmt.Println("sum of even numbers =", s2)

	s3 := sumOdd(sumI, ii...)
	fmt.Println("sum of odd numbers =", s3)
}

func sumEven(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func sumOdd(f func(xi ...int) int, vi ...int) int {
	var oi []int
	for _, v := range vi {
		if v%2 != 0 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}
