package main

import "fmt"

func main() {
	i := 633
	fmt.Printf("%d\t%b\t%#x\n", i, i, i)
	i2 := i << 1
	fmt.Printf("%d\t%b\t%#x\n", i2, i2, i2)

}
