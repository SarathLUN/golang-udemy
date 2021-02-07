package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			fmt.Println("goroutine:", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)

}

// run with ->go run -race 03-exercise-03.go
