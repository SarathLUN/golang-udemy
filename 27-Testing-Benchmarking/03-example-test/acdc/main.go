// Package acdc: are you ready to rock!
package acdc

// Sum: take variadic parameters and return the sum
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
