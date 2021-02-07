package main

import "fmt"

func main() {
	x := 46
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T\n", x)  // int is a type
	fmt.Printf("%T\n", &x) // *int is another difference type

	b := &x
	fmt.Println(b)
	fmt.Println(*b)

	*b = 99
	fmt.Println(x)

}
