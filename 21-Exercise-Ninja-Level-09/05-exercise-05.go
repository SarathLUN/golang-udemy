package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var incrementer int64

	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			fmt.Println("goroutine:", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

// run with ->go run -race 03-exercise-03.go
