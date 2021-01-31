package main

import "fmt"

func main() {
	c := make(<-chan int, 1)
	//c <- 43 //this will be error
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}
