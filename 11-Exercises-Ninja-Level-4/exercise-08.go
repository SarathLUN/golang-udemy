package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)
	for k, v1 := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v1 {
			fmt.Printf("\t index-%d = %v \n", i, v2)
		}
	}
}
