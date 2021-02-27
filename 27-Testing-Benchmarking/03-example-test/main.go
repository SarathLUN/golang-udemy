package main

import (
	"fmt"
	"golang-udemy/27-Testing-Benchmarking/03-example-test/acdc"
)

func main() {
	fmt.Println(acdc.Sum(1, 2, 3, 4))
	fmt.Println(acdc.Sum(1, 2, -1))
}
