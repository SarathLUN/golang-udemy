package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		if x == 5{
			break
		}
		fmt.Println(x)
		x++
	}
}
