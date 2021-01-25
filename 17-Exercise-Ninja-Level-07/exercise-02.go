package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(xP *person) {
	xP.age++
	fmt.Println(&xP) // get mem address 0xc00000e030
	fmt.Println(xP)  // get &{Tony 46}
	fmt.Println(*xP) // get struct {Tony 46}
	(*xP).age++
}

func main() {
	p1 := person{
		"Tony",
		45,
	}
	fmt.Println(p1)
	fmt.Println("before: ", &p1)     //actually there is no address for a struct
	fmt.Println("before: ", &p1.age) //but there is an address for value of the struct

	changeMe(&p1)
	fmt.Println(p1)
}
