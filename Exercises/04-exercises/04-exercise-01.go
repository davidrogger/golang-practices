/**
	- Usando uma literal composta:
		- Crie um array que suporte 5 valores do tipo int
		- Atribua valores aos seus índices
	- Utilize range e demonstre os valores do array.
	- Utilizando format printing, demostre o tipo do array.
**/

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, value := range array {
		fmt.Printf("O tipo de %v é %T\n", value, value)
	}
}
