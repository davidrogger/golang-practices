/*
	- Crie um struct "pessoa"
	- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
	Essa função deve mudar o valor armazenado no endereço *pessoa.
		- Dica: a maneira "correta" para fazer dereferencia de um valor em struc seria
		(*valor).campo
		- Mas consta um exceção na documentação.
		link: https://golang.org/ref/spec#Selectors
		- "As an exception, if the type of x is a named pointer type and
		(*x).f is a valid selector expression denoting a field (but not a method),
		x.f is shorthand for (*x).f."
		- Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
	- Crie um valor do tipo pessoa;
	- Use a função mudeMe e demonstre o resultado
*/

package main

import "fmt"

type pessoa struct {
	nome string
}

func mudeMe(p *pessoa) {
	(*p).nome = "bety"
}

func main() {
	jonas := pessoa{"jonas"}
	fmt.Println(jonas)
	mudeMe(&jonas)
	fmt.Println(jonas)
}
