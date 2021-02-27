package acdc

import "fmt"

func ExampleSum_first() {
	fmt.Println(Sum(2, 3))
	// Output: 5
}
func ExampleSum_second() {
	fmt.Println(Sum(2, 3, 4))
	// Output:
	// 9
}
