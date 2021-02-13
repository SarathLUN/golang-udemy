package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("from defer func")
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened:", err)
		//log.Println("err log:", err)
		panic(err)
		//log.Fatalln(err)
	}
}
