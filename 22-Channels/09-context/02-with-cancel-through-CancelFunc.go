package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:", ctx)
	fmt.Println("context error:", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("---")

	ctx, _ = context.WithCancel(ctx)
	/*warning:
	the cancel function returned by context.WithCancel should be called,
	not discarded, to avoid a context leak
	*/

	fmt.Println("context:", ctx)
	fmt.Println("context error:", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("---")
}
