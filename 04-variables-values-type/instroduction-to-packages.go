package main

import "fmt"

func main() {
	//n, err := fmt.Println("Hello, Playground", 42.0, true)  // if we catch the error, we must use this var somewhere, otherwise this will error
	//fmt.Println(err)
	n, _ := fmt.Println("Hello, Playground", 42.0, true) // "_" will through the error away and return
	fmt.Println(n)

}
