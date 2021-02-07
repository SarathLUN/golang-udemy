package main

import "fmt"

var x, y int // declare in 1 line for same type

func main() {
	x = 43
	y = 66
	fmt.Println(x, y)

	x, y = y, x // swap value of x & y
	fmt.Println(x, y)

	x, y = 88, 99 // re-assign in 1 line
	fmt.Println(x, y)
}
