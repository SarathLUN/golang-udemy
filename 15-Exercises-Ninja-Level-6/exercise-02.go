package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	r := foo2(x...)
	fmt.Println(r)

	r2 := bar2(x)
	fmt.Println(r2)
}

func foo2(v1 ...int) int {
	s := 0
	for _, v := range v1 {
		s += v
	}
	return s
}

func bar2(p1 []int) int {
	s := 0
	for _, v := range p1 {
		s += v
	}
	return s
}
