/*
	- Escreva um programa que coloque 100 números em um canal, retire os números do
	canal e demonstre-os
*/

package main

import "fmt"

func main() {
	canal := gen()
	receive(canal)
}

func gen() chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
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
