package main

import "fmt"

func main() {
	c := make(chan int)
	// do send
	go foo(c)
	// do receive
	bar(c)
	fmt.Println("about to exit.")
}

func foo(c chan<- int) {
	c <- 43
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
