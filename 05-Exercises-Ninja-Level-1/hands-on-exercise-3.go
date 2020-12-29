package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
	s := fmt.Sprintf("%v\t%v\t%v",x,y,z)

	// b. print out the value stored by variable “s”
	fmt.Println(s)
}
