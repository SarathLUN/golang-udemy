package main

import "fmt"

func main() {
	fmt.Printf("true && true\t becom %v\n", true && true)
	fmt.Printf("true && false\t becom %v\n", true && false)
	fmt.Printf("true || true\t becom %v\n", true || true)
	fmt.Printf("true || false\t becom %v\n", true || false)
	fmt.Printf("!true\t\t becom %v\n", !true)
	fmt.Printf("!false\t\t becom %v\n", !false)
}
