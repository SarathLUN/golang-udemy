package main

import "fmt"

func main() {
	xi := []int{2, 3, 5, 6}
	x := sum(xi...)
	fmt.Println("x = ", x)
	fmt.Println("---")

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
