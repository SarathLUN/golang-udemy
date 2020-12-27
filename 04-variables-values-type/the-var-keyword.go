package main

import "fmt"

var z = "hi, i am out side of func"
func main() {
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("again: ",z)
}

// key take away:
// 1. var can use at the global scope, but short declaration can use only in side the function