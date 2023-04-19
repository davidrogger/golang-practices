/*
	- Crie e utilize uma função anonima
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Funcao anonima")
	}()
}
