package main

import "fmt"

func bar4() func() int {
	return func() int {
		return 64
	}
}

func main() {
	fmt.Println(bar4()())
}
