package main

import "fmt"

//For this exercise:
//1. Create your own type. Have the underlying type be an int.
type my_Int int
func main() {
	//2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword in func main:
	var x my_Int
	//		a. print out the value of the variable “x”
	fmt.Println(x)
	//		b. print out the type of the variable “x”
	fmt.Printf("%T\n",x)
	//		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42
	//		d. print out the value of the variable “x”
	fmt.Println(x)
}

