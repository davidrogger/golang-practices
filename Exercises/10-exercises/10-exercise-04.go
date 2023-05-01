/*
	- Utilizando este código: https://go.dev/play/p/MvL6uamrJP
	- Use um select statement para demonstrar os valores do canal.
*/

package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive2(c, q)

	fmt.Println("about to exit")
}

func gen(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		case v, ok := <-q:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		}
	}
}

func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
