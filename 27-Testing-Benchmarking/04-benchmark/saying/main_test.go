package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"Tony1"}, "Welcome my dear, Tony1"},
		{"test2", args{"Tony2"}, "Welcome my dear, Tony2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Tony"))
	// Output:
	// Welcome my dear, Tony
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Tony")
	}
}
