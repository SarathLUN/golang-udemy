package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Errorf("got %v, want 70", n)
	}
}

func TestYearTwo(t *testing.T) {
	n := YearTwo(10)
	if n != 70 {
		t.Errorf("got %v, want 70", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearTwo() {
	fmt.Println(YearTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearTwo(10)
	}
}
