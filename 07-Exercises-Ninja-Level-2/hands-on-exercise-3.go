package main

import "fmt"

const (
	type_const   string = "Hello, World!"
	untype_const        = "Hello, Another World!"
)

func main() {
	fmt.Println(type_const)
	fmt.Println(untype_const)
}
