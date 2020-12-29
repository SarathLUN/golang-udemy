package main

import "fmt"

var y = 43
var a string = `James said, 
"Shaken, 

not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n",y)

	z := "Shaken, not stirred"
	fmt.Printf("%T\n",z)

	fmt.Println(a)
}
