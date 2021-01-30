package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println(p.name, ", hello!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Jame"}
	saySomething(&p1)
}
