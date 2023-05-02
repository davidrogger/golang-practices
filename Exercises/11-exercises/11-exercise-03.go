/*
	- Crie um struct "erroEspecial" que implemente a interface builtin.error
	- Crie uma função que tenha um valor do tipo error como parâmetro.
	- Crie um valor do tipo "errorEspecial" e passe-o para a função da
	instrução anterior.
*/

package main

import (
	"fmt"
)

type erroEspecial struct {
	myError string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("especial: %v", e.myError)
}

func funcao(err error) {
	fmt.Println(err.(erroEspecial).myError)
}

func main() {
	e := erroEspecial{"teste"}
	funcao(e)
}
