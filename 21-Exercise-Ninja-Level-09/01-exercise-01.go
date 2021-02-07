package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("foo:", i)
		}
		fmt.Println("goroutine:", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		for i := 5; i < 10; i++ {
			fmt.Println("bar:", i)
		}
		fmt.Println("goroutine:", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main run")
	fmt.Println("goroutine:", runtime.NumGoroutine())
}
