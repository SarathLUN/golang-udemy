package main

import "fmt"

func main() {
	c := make(chan<- int, 1)
	c <- 43
	//fmt.Println(<-c) //this will be error
	fmt.Printf("%T\n", c)
}
