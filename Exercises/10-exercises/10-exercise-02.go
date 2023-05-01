/*
	- Faça esse código funcionar: https://go.dev/play/p/oB-p3KMiH6
*/

package main

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("-------\n")
	fmt.Printf("cs\t%T\n", cs)
}
