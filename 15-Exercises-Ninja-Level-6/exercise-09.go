package main

import "fmt"

func foo4(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func main() {
	g := func(xi []int) int {
		return xi[0] + xi[len(xi)-1]
	}
	f := foo4(g, []int{1, 2, 3})
	fmt.Println(f)
}
