/*
Package main of Exercise Ninja Level 11 and hand-on exercise number 4
*/
package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v, %v, %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("ErrFromSqrt: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
