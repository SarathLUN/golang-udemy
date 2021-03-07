package main

import (
	"fmt"
	"golang-udemy/28-Exercise-Ninja-Level-13/exercise-01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		"Fido",
		dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearTwo(20))
}
