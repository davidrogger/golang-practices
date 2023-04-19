/*
	- Demonstre o funcionamento de closure
	- Crie uma função que retornar outra função, onde esta outra função
	faz uso de uma variável além de seu scope
*/

package main

import "fmt"

func main() {
	f := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}

	}

	a := f()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := f()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}
