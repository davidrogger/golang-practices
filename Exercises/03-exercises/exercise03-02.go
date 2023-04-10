// Imprima na tela: O unicode code point de todas as letras mai´sculas do alfabeto, três vezes cada.
// Exemplo:
// 65
// 	U+0041 'A'
// 	U+0041 'A'
// 	U+0041 'A'
// 66
// 	U+0042 'B'
// 	U+0042 'B'
// 	U+0042 'B'

package main

import "fmt"

func main() {
	firstLetter := 65
	lastLetter := 90

	for l := firstLetter; l <= lastLetter; l++ {
		fmt.Println(l)
		for s := 1; s <= 3; s++ {
			fmt.Printf("\t%#U\n", l)
		}
	}
}
