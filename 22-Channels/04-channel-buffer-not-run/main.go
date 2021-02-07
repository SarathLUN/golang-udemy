package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 44
	c <- 46 //channel buffer is overloaded
	fmt.Println(<-c)
}
