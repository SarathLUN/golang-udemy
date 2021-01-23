package main

import "fmt"

func main() {
	defer foo2() // will execute at the last of function
	defer bar3() // defer will run back in order(last in first out)
	bar2()
}

func foo2() {
	fmt.Println("foo2")
}

func bar2() {
	fmt.Println("bar2")
}

func bar3() {
	fmt.Println("bar3")
}
