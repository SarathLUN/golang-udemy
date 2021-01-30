package main

import "fmt"

func time5(x int) int {
	return x * 5
}

func main() {
	//create channel
	ch := make(chan int)
	// goroutine on anonymous func and grape result into channel
	// in this situation without channel we will not get value from goroutine
	go func() {
		ch <- time5(5)
	}()
	// read value from channel
	fmt.Println(<-ch)
}
