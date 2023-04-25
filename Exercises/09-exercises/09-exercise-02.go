/*
	- Crie um tipo para um struct chamado "pessoa"
	- Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
	- Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
	- Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame
	o método "falar"
	- Demonstre no seu código:
		- Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
		- Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/

package main

import "fmt"

type pessoa struct {
	name string
	age  int
}

type humanos interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.name, "e tenho", p.age, "anos.")
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	jonas := pessoa{"Jonas", 33}

	dizerAlgumaCoisa(&jonas)
}
