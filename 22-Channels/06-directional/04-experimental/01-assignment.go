package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive only
	cs := make(chan<- int) // send only

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	/*c = cr
	c = cs*/

	/* output:
	# command-line-arguments
	./01-assignment.go:16:4: cannot use cr (type <-chan int) as type chan int in assignment
	./01-assignment.go:17:4: cannot use cs (type chan<- int) as type chan int in assignment
	*/

	// specific to specific doesn't assign
	//cs = cr

	/*output:
	# command-line-arguments
	./01-assignment.go:26:5: cannot use cr (type <-chan int) as type chan<- int in assignment
	*/

	// general to specific assigns
	/*	cr = c
		cs = c

		fmt.Println("-----")
		fmt.Printf("c\t%T\n", c)
		fmt.Printf("cr\t%T\n", cr)
		fmt.Printf("cs\t%T\n", cs)

		fmt.Println("-----")
		fmt.Println(<-c)
		fmt.Println(<-cr)
		fmt.Println(<-cs) // this also error, look like it not work together at all
	*/

}
