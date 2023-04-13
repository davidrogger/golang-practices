/**
	- Crie uma slice usando make que possa conter todos os estados do Brasil.
		- Os estados: "Acre", "Alagados", "Amapá", "Amazonas", "Bahia", "Cerá", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande so Sul", "Rondônia", "Roraima", "Santa Caterina", "São Paulo", "Sergipe", "Tocantins".
	- Demonstre o len e cap da slice.
	- Demonstre todos os valores da slice
	- Sem utilizar o range.
**/

package main

import "fmt"

func main() {
	c := make([]string, 26, 26)
	c = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(cap(c))
	fmt.Println(len(c))
	fmt.Println(c)
}
