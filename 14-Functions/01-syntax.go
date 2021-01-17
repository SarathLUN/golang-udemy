package main

import "fmt"

func main() {
	foo()
	bar("James")
	result := woo(1, 2)
	fmt.Println(result)
}

func woo(i int, j int) int {
	return i + j
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func foo() {
	fmt.Println("hello from foo.")
}
