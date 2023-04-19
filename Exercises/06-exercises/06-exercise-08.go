/*
	- Crie uma função que retorna uma função
	- Atribua a função retornada a uma variavel
	- Chame a função retornada
*/

package main

import "fmt"

func callFunc(n1 int) func(int) int {
	return func(n2 int) int {
		return n1 + n2
	}
}

func main() {
	base := callFunc(5)
	total := base(5)

	fmt.Println(total)
}
