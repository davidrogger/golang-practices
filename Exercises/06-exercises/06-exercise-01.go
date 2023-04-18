/*
	- Crie uma função que retorne um int
	- Outra que retorne um int e uma string
	- Demonstre seus resultados
*/

package main

import "fmt"

func getNumber(n int) int {
	return n
}

func getNumberString(n int, s string) (int, string) {
	return n, s
}

func main() {
	n := getNumber(1)
	n2, s := getNumberString(2, "jonas")
	fmt.Println(n, n2, s)
}
