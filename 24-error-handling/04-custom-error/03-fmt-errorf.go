package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	x, err := sqrt(-10.5)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(x)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	return math.Sqrt(f), nil
}
