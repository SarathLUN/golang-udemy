package main

import "fmt"

func main() {
	// composite literal
	x := []int{3, 4, 6}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
