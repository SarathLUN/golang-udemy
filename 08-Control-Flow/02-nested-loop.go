package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i = %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("\tj = %d\n", j)
		}
	}
}
