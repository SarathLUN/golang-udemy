package main

import "fmt"

func foo3() func() int {
	return func() int {
		return 43
	}
}

func main() {
	f := foo3()
	fmt.Println(f())
}
