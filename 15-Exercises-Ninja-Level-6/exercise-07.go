package main

import "fmt"

func main() {
	f := func(x int) {
		for i := 0; i < x; i++ {
			fmt.Println(i)
		}
	}
	f(6)
}
