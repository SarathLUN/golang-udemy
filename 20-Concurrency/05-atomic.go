package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	var counter int64 //declare int64 for using Atomic

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) //add counter to atomic
			runtime.Gosched()
			fmt.Println("counter:", atomic.LoadInt64(&counter)) //load counter from atomic
			wg.Done()
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
