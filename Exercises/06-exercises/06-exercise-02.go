/*
	- [x] Crie uma função que receba um parâmetro variádico do tipo int e
	retorne a soma de todos os ints recebidos;
	- [x] Passe um valor do tipo slice of int com argumento para função;
	- [x] Crie outra função, esta deve receber um valor do tipo slice of int e retornar
	a soma de todos os elementos da slice;
	- [x] Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func sumOfNumbers(n ...int) int {
	total := 0
	for _, number := range n {
		total += number
	}

	return total
}

func sumOfSlice(n []int) int {
	return sumOfNumbers(n...)
}

func main() {
	numbers := []int{1, 1, 1, 1, 1}
	total := sumOfNumbers(numbers...)
	total2 := sumOfSlice(numbers)

	fmt.Println(total, total2)
}
