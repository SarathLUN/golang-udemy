package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func (n ByName) Len() int {
	return len(n)
}
func (n ByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n ByName) Less(i, j int) bool {
	return n[i].Name < n[j].Name
}

func main() {
	p1 := Person{"Jame", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}
	people := []Person{p1, p3, p2, p4}
	fmt.Println(people)

	// to be able to use sort.Sort(data Interface), we need type sort.Interface
	// let convert people to be type ByAge, which eq to Interface
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
