/*
	- Para testar sua habilidade de se virar por conta propria... desafio:
		- Utilizando as seguintes fontes:
			https://pkg.go.dev/testing?utm_source=godoc
			https://www.golang-book.com/books/intro/12
		- Tente descobrir por conta prórpia como funcionam teste em Go.
		- O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.
*/

package myDiv

import "testing"

func TestMyDiv(t *testing.T) {
	result := myDiv(1, 2)
	expect := 0.5
	if result != expect {
		t.Error("Expected:", expect, "Got:", result)
	}
}
