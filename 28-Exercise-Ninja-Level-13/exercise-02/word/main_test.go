package word

import (
	"fmt"
	"github.com/SarathLUN/golang-udemy/28-Exercise-Ninja-Level-13/exercise-02/quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Errorf("want: 1, got: %v\n", v)
			}
		case "two":
			if v != 1 {
				t.Errorf("want: 1, got: %v\n", v)
			}
		case "three":
			if v != 3 {
				t.Errorf("want: 3, got: %v\n", v)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Errorf("want: 3, got: %v \n", n)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}