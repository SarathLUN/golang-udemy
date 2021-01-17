package main

import "fmt"

func main() {
	//basic func
	foo()

	//takes an argument
	bar("James")

	//return
	result := woo(1, 2)
	fmt.Println(result)

	//multiple returns
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
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
