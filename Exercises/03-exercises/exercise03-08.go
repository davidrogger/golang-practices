// Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {
	numb := 11
	switch {
	case numb > 20:
		fmt.Println("Número maior que 20")
	case numb > 10:
		fmt.Println("Número maior que 10")
	default:
		fmt.Println("Menor que 20 e 10")
	}
}
