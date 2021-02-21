package main

import (
	"fmt"
	dog "golang-udemy/26-Exercise-Ninja-Level-12/01-exercise-01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
