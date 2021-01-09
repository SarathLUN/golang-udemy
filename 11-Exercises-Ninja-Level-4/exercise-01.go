package main

import "fmt"

func main() {
	//create an ARRAY which holds 5 VALUES of TYPE int
	a := [5]int{}
	fmt.Println(a)
	//assign VALUES to each index position.
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	fmt.Println(a)
	//Range over the array and print the values out.
	for _, v := range a {
		fmt.Println(v)
	}
	//print out the TYPE of the array
	fmt.Printf("%T\n", a)
}
