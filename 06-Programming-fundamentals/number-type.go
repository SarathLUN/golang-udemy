package main

import "fmt"

var x2 int64
var y float32

func main() {
	x2 = 43
	y = 43.5895764
	fmt.Println(x2)
	fmt.Println(y)
	fmt.Printf("%T\n", x2)
	fmt.Printf("%T\n", y)
}
