package main

import "fmt"

// un-type constant
const (
	a = 64
	b = 43.46
	c = "Tony"
)

// typed constant
const (
	d int     = 87
	e float64 = 98.46
	f string  = "Hello world."
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
