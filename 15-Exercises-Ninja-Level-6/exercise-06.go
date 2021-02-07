package main

import "fmt"

func main() {
	func(x int) {
		for i := 0; i < x; i++ {
			fmt.Println(i)
		}
		fmt.Println("done!")
	}(5)

}
