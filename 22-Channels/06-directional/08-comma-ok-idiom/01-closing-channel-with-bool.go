package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//quit <- true  // this will show the difference
	close(even)
	close(odd)
	close(quit) // without line#28 this channel is closed with empty
}

// receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		/*case i, ok := <-quit:
		if !ok {
			fmt.Println("can't access channel(quit):i=", i, "ok=", ok)
		} else {
			fmt.Println("successful access channel(quit), i=", i, "ok=", ok)
		}
		return*/
		case i := <-quit:
			fmt.Println("i=", i)
			return
		}
	}
}

/*Note: https://golang.org/ref/spec#Receive_operator
yields an additional untyped boolean result reporting whether the communication succeeded.
The value of ok is true if the value received was delivered by a successful send operation to the channel,
or false if it is a zero value generated because the channel is closed and empty.

=> so we should always use ", ok" to check availability of channel
however, Goland provide "zero value" which still work well event we don't check
but we don't know the value is really retrieve from channel or not
*/
