package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	incrementer := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			//mu.Unlock() //if we put here, we still get data race, because below code is trying to read incrementer
			fmt.Println(incrementer)
			mu.Unlock()
			fmt.Println("goroutine:", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

// run with ->go run -race 03-exercise-03.go
