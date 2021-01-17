package main

import "fmt"

func main() {
	x := sum(2, 3, 5, 6)
	fmt.Println("X = ", x)
	fmt.Println("---")

	y := sum(5, 2, 7, 9, 10, 334, 8, 4, 5)
	fmt.Println("y = ", y)
}

func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	sum := 0
	for i, v := range i {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("Sum = ", sum)
	return sum
}
