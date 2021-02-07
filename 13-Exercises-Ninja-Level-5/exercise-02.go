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

	m := map[string]person{
		p1.lastName:p1,
		p2.lastName:p2,
	}
	fmt.Println(m)
	for k,v := range m{
		fmt.Println(k)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _,val := range v.favoriteFlavors{
			fmt.Printf("\t %v \n",val)
		}
		fmt.Println("----")
	}

}
