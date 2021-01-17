package main

import "fmt"

func main() {
	foo()
	bar("James")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func foo() {
	fmt.Println("hello from foo.")
}
