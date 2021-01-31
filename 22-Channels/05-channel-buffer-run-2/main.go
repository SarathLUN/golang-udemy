package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 44
	c <- 46
	fmt.Println(<-c) //FIFO will return 44
	fmt.Println(<-c) //FIFO will return 46

	/*fmt.Println("--- run experiment1 ---")
	experiment1()*/

	/*fmt.Println("--- run experiment2 ---")
	experiment2()*/
}

func experiment1() {
	c1 := make(chan int)
	c1 <- 1
	a := <-c1
	c1 <- 2
	b := <-c1
	fmt.Println(a, b)
}

func experiment2() {
	c2 := make(chan int)
	fmt.Println(<-c2)

}
