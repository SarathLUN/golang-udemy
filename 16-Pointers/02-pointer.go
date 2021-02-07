package main

import "fmt"

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 88
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}

func main() {
	x := 0
	fmt.Println("x before", x)
	fmt.Println("x before", &x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}
