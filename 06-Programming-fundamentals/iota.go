package main

import "fmt"

const (
	a2 = iota
	b2
	c2
)

const (
	d2 = iota
	e2
	f2
)

func main() {
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)
	fmt.Println(e2)
	fmt.Println(f2)
}
