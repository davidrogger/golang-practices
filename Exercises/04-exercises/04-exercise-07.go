/**
	- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
	- "Nome", "Sobrenome", "Hobby Favorito"
	- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
**/

package main

import "fmt"

func main() {
	persons := [][]string{
		[]string{"Nome", "Sobrenome", "Hobby Favorito"},
		[]string{"jonas", "Doe", "Mock"},
		[]string{"modiz", "beserra", "dirty"},
		[]string{"biel", "hate", "play"},
	}
	for _, person := range persons {
		name, surname, hobby := person[0], person[1], person[2]
		fmt.Println(name, surname, hobby)
	}
}
