/*
	- [x] Crie um tipo struct "pessoa" que contenha os campos:
		- nome
		- sobrenome
		- idade
	- [x] Crie um método para "pessoa" que demonstre o nome completo e a idade;
	- [x] Crie um valor de tipo "pessoa"
	- [x] Utilize o método criado para demonstrar esse valor
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) showInfo() {
	fullName := p.nome + " " + p.sobrenome
	fmt.Printf("Nome completo: %v, idade: %v\n", fullName, p.idade)
}

func main() {
	jonas := pessoa{"Jonas", "Doe", 33}
	jonas.showInfo()
}
