package main

import "fmt"

//at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x2”

type hotdog int

var x2 hotdog

//var y int // experiment#1

func main() {
	//print out the value of the variable “x2”
	fmt.Println(x2)
	//print out the type of the variable “x2”
	fmt.Printf("%T\n", x2)
	//assign your own VALUE to the VARIABLE “x2” using the “=” OPERATOR
	x2 = 46
	//print out the value of the variable “x2”
	fmt.Println(x2)
	//now use CONVERSION to convert the TYPE of the VALUE stored in “x2” to the UNDERLYING TYPE
	//y = int(x2) // experiment#1
	y := int(x2)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
