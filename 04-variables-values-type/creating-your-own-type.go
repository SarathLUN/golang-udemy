package main

import "fmt"

var i int

type hotdog int
var h hotdog

func main() {
	i = 46
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	h = 66
	fmt.Printf("%T\n", h)
}
