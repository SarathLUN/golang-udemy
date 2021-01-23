package main

import (
	"fmt"
	"log"
	"os/exec"
)

type person struct {
	first  string
	last   string
	gender bool
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	voice := "can"
	if !s.ltk {
		voice = "can't"
	}
	sx := "I am " + s.first + s.last + " and I " + voice + " kill."
	fmt.Println(sx)
	cmd := exec.Command("espeak", sx) //https://golangr.com/text-to-speech/
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			true,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
			false,
		},
		ltk: false,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
