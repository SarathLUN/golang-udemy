// package word provide custom function for working with strings and words.
package word

import "strings"

// UseCount return the number of times each word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return a number of words in a string.
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
