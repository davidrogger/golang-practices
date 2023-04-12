// Crie um programa que utilize a declaração switch, onde switch statement seja uma variável do tipo string e identificador "esporteFavorito"

package main

import "fmt"

func main() {
	esporteFavorito := "a"

	switch esporteFavorito {
	case "basquete":
		fmt.Println("Esporte favorito é basquete")
	case "futebol":
		fmt.Println("Esporte favorito é futebol")
	case "volei":
		fmt.Println("Esporte favorito é volei")
	default:
		fmt.Println("Nenhum esporte favorito")
	}
}
