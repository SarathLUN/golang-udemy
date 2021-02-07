package main

import "fmt"

func main() {
	xss := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	for i, xs := range xss {
		fmt.Println("record:", i)
		for j, v := range xs {
			fmt.Printf("\t index-%d = %v \n", j, v)
		}
	}

}
