package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	m["Tommy"] = 33
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, "=", v)
	}

	// Map - delete
	delete(m, "James")
	fmt.Println(m)
	// we also can delete not existed element without error
	delete(m, "Hello")
	fmt.Println(m)
}
