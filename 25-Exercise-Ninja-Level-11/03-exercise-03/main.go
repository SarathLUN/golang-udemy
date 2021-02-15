package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("here is erro: %v", c.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
	fmt.Println("foo ran with assertion:", e.(customErr).info)

}
