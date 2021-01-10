package main

import "fmt"

// please ignore the inspection error for now on duplicated declaration
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)

	// fields in person has promoted to secretAgent, thus we no need to use sa1.person.first
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
}
