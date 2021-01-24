package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	i, s := bar() // this is how to catch multiple return values
	fmt.Println(i)
	fmt.Println(s)
}

func foo() int {
	return 46
}

func bar() (int, string) {
	return 64, "Woooh"
}
