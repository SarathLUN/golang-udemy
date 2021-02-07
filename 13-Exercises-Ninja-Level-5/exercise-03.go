package main

import "fmt"

type vehicle struct {
	door int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle{
			door:  2,
			color: "white",
		},
		true,
	}
	fmt.Println(t)

	s:= sedan{
		vehicle{
			door:  4,
			color: "black",
		},
		false,
	}
	fmt.Println(s)

	fmt.Println(t.door)
	fmt.Println(s.color)
}