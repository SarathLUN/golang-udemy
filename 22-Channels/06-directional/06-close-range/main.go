package main

import "fmt"

func main() {
	c := make(chan int)
	// loop sending
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close(c)
	}()
	// receive from looping channel
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit.")
	/*Note:
	if we try to range over un-closed channel it will through an error:
	---
	0
	1
	2
	3
	4
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan receive]:
	main.main()
		/Users/tony/go-space/src/golang-udemy/22-Channels/06-directional/06-close-range/main.go:15 +0xe5
	exit status 2
	---
	*/
}
