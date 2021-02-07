package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Anonymous func with parameter x =", x)
	}(46)
}
