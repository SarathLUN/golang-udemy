package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c *circle) area2() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}
type shape2 interface {
	area2() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func info2(s shape2) {
	fmt.Println("area", s.area2())
}

func main() {
	c := circle{5}
	info(c)   //NON-POINTER RECEIVER & NON-POINTER VALUE
	info(&c)  //NON-POINTER RECEIVER & POINTER VALUE
	info2(&c) //POINTER RECEIVER & POINTER VALUE
	//info2(c) //POINTER RECEIVER & NON-POINTER VALUE -> ERROR
	fmt.Println(c.area2())
}
