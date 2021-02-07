package main

import "fmt"

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan int) // we need q to return from select statement

	// send
	go send(e, o, q)

	// receive
	receive(e, o, q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("even channel:", v)
		case v := <-o:
			fmt.Println("odd channel:", v)
		case v := <-q:
			fmt.Println("quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	// best practise is to close sent channel after all
	close(e)
	close(o)
	close(q)
}
