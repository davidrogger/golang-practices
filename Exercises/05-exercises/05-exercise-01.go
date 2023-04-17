/**
	- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
		- nome
		- sobrenome
		- Sabores favoritos de sorvete
	- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
**/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	jonas := pessoa{"jonas", "doe", []string{"chocolate", "napolitano"}}
	bahdok := pessoa{"bahdok", "krowe", []string{"morango", "pistache"}}

	clients := []pessoa{jonas, bahdok}

	for _, client := range clients {
		fmt.Printf("%v %v\n", client.nome, client.sobrenome)
		for index, favor := range client.sabores {
			fmt.Println(index+1, "-", favor)
		}
	}
}
