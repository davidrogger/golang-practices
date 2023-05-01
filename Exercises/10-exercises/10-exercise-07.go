/*
	- Crie um programa que lance 10 goroutines onde cada um envia 10 números
	a um canal;
	- Tire estes números do canal e demonstre-os
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int)

	go multiRoutine(10, channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func multiRoutine(quantity int, c chan int) {
	for i := 1; i <= quantity; i++ {
		wg.Add(1)
		go send(100, c)
	}
	wg.Wait()
	close(c)
}

func send(quantity int, c chan int) {
	for i := 1; i <= quantity; i++ {
		c <- i
		fakeWorkload()
	}
	wg.Done()
}

func fakeWorkload() {
	time.Sleep(time.Millisecond * 1000)
}
