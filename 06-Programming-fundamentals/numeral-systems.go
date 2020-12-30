package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	nbs := bs[0]
	fmt.Println(nbs)
	fmt.Printf("%T\n", nbs)
	fmt.Printf("%b\n", nbs)
	fmt.Printf("%#x\n", nbs)
}
