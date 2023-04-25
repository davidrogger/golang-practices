/*
	- Além da goroutine principal, crie duas outras goroutines
	- Cada goroutine adicional devem fazer um print separado.
	- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o
	programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go func() {
		fmt.Println("Primeiro Go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Segunda Go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Terceira Go routine")
		wg.Done()
	}()

	wg.Wait()
}
