/*
	- Aplicando mutex
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	const routineAmount = 100
	wg.Add(routineAmount)

	for i := 1; i <= routineAmount; i++ {
		go func() {
			mutex.Lock()
			runtime.Gosched()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}

// go run -race Exercises/09-exercises/09-exercise-04.go
