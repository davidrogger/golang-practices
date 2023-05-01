/*
	- Utilizando este código: https://go.dev/play/p/sfyu4Is3FG
	... Use um for range loop para demonstrar os valores do canal.
*/

package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("About to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
