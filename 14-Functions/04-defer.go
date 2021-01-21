package main

import "fmt"

func main() {
	defer foo2() // will execute at the last of function
	bar2()
}

func foo2() {
	fmt.Println("foo2")
}

func bar2() {
	fmt.Println("bar2")
}
