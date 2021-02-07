package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "lang", "Go")
	findValue(ctx, "lang")
	//add more value
	ctx = context.WithValue(ctx, "dog", "")
	findValue(ctx, "dog")
	findValue(ctx, "color")
}

func findValue(ctx context.Context, key string) {
	if v := ctx.Value(key); v != nil && v != "" {
		//fmt.Println("found value:",v,", for key:",key)
		fmt.Printf("%v : %v\n", key, v)
		return
	}
	fmt.Println("key not found:", key)
}
