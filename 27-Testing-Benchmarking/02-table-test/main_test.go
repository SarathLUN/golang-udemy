package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{2, 3, 4}, 9},
		{[]int{1, 2}, 3},
		{[]int{1, -2, 10, 9}, 9}, //not really the answer but just want to show 1 error
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Errorf("Expected %v, Got: %v", v.answer, x)
		}
	}
}
