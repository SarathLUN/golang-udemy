package main

import "fmt"

func main() {
	x := []int{3, 4, 6, 34, 46}
	fmt.Println(x)

	x = append(x, 123, 456)
	fmt.Println(x)

	y := []int{77, 88, 99}
	x = append(x, y...) // "..." = unpack sth
	fmt.Println(x)

	// delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
