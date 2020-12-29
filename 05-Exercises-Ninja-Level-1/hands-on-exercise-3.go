package main

import "fmt"

var i = 42
var s = "James Bond"
var b = true

func main() {
	// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
	r := fmt.Sprintf("%v\t%v\t%v", i, s, b)

	// b. print out the value stored by variable “s”
	fmt.Println(r)
}
