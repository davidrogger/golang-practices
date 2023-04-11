// Imprima o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fourRest := i % 4
		fmt.Printf("O resto da divisão de %v e 4 é %v\n", i, fourRest)
	}
}
