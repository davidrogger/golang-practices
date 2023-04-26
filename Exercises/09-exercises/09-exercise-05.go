/*
	- Aplicando atomic
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func main() {
	const routineAmount = 100
	wg.Add(routineAmount)

	for i := 1; i <= routineAmount; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}

// go run -race Exercises/09-exercises/09-exercise-05.go
