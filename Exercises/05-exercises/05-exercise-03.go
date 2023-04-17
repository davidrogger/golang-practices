/**
	- Crie um novo tipo: véiculo
		- O tipo subjacente deve ser struct
		- Deve conter os campos: porta, cor
	- Crie dois tipos, caminhonete e sedan
		- Os tipos subjacentes deve ser struct
		- Ambos devem conter "veículo" como struct embutido
		- O tipo caminhonete deve conter um campo bool chamado "quatrorodas"
		- O tipo sedan deve conter um campo bool chamado "modeloLuxo"
	- Usando os structs véiculos, caminhonete e sedan:
		- Usando composite literal, crie um valor de tipo caminhonete e
		dê valores a seus campos.
		- Usando camposite literal, crie um valor de tipo sedan e dê valores
		a seus campos.
	- Demonstre estes valores.
	- Demonstre um único campo de cada um dos dois.
**/

package main

import "fmt"

type veiculo struct {
	porta int
	cor   string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	caminhoneteNova := caminhonete{veiculo{2, "azul"}, true}
	sedanNovo := sedan{veiculo{4, "cinza"}, false}

	fmt.Println(caminhoneteNova, sedanNovo)
}
