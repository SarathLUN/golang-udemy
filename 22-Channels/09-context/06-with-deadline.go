package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(5 * time.Second)
	fmt.Printf("%T\n", deadline)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("context error:", ctx.Err())
	}
}
