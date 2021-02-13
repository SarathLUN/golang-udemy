package main

import "fmt"

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("char:", n)
}
