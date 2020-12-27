package main

import "fmt"

var b = 43

func main() {
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	fmt.Printf("%b\n",b)
	fmt.Printf("%x\n",b)
	fmt.Printf("%#x\n",b)

	y := 46
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v", y)
}
