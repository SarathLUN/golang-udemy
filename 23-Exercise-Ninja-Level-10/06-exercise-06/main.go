package main

import "fmt"

func main() {
	c := make(chan int)
	// send 100 numbers to the channel
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	// receive 100 numbers from the channel
	for v := range c {
		fmt.Println(v)
	}
}
