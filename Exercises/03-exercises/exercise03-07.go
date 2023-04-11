// Crie um programa que demonstre o funcionamento da declaração if e else

package main

import "fmt"

func main() {
	x := 0
	if x > 10 {
		fmt.Println("Maior que 10")
	} else if x >= 1 {
		fmt.Println("Maior que 1")
	} else {
		fmt.Println("Menor que 1 e 10")
	}
}
