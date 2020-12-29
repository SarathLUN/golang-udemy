package main

import "fmt"

func main() {
	x := 42 // declare and assign
	fmt.Println(x)
	x = 99 // re-assign
	fmt.Println(x)
	//x = "hello"  // Error: cannot use "hello" (type untyped string) as type int in assignment
	//fmt.Println(x)

	y := 100 + 24
	fmt.Println(y)
}