/**
	- Utilizando a solução anterior, coloque os valores do tipo "pessoa"
	num map, utilizando os sobrenomes como key.
	- Demonstre os valores do map utilizando range.
	- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro
	do range anterior.
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

	clients := map[string]pessoa{"doe": jonas, "krowe": bahdok}

	for _, client := range clients {
		fmt.Printf("%v %v\n", client.nome, client.sobrenome)
		for index, favor := range client.sabores {
			fmt.Println(index+1, "-", favor)
		}
	}
}
