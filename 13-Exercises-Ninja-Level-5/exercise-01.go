package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		firstName: "Tony",
		lastName:  "Stark",
		favoriteFlavors: []string{
			"vanilla",
			"chocolate",
			"coconut",
		},
	}
	fmt.Println(p1)

	p2 := person{
		firstName: "Doe",
		lastName:  "John",
		favoriteFlavors: []string{
			"chocolate",
			"strawberry",
		},
	}
	fmt.Println(p2)

	fmt.Println("-----")
	fmt.Printf("%v.%v likes:\n", p1.firstName, p1.lastName)
	for _, v := range p1.favoriteFlavors {
		fmt.Printf("\t- %v\n", v)
	}

	fmt.Println("-----")
	fmt.Printf("%v.%v likes:\n", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteFlavors {
		fmt.Printf("\t- %v\n", v)
	}
}
