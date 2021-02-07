package main

import "fmt"

func main() {
	x := 80
	if x < 60 {
		fmt.Println("Grade: F")
	} else if x == 60 && x < 70 {
		fmt.Println("Grade: E")
	} else if x < 80 {
		fmt.Println("Grade: D")
	} else if x < 90 {
		fmt.Println("Grade: C")
	} else if x < 95 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: A")
	}
}
