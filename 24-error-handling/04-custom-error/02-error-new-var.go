package main

import (
	"errors"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	x, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(x)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
