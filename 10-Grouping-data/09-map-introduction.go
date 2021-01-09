package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barn"]) // this will return zero

	// how we check key is exist, call ", ok" comma ok idiom
	v, ok := m["Barn"]
	fmt.Println(v)
	fmt.Println(ok)

	// so we can make it as combine idiomatic if statement
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}
}
