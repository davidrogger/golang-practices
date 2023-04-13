/**
	- Crie um map com key tipo string e value tipo []string
		- Key deve conter nomes no formato sobrenome_nome
		- Value deve conter os hobbies favoritos da pessoa
	- Demonstre todos esses valores e seus índices.
**/

package main

import "fmt"

func main() {
	users := map[string][]string{
		"doe_Jonas": []string{"Mock", "testing"},
	}

	for name, hobbies := range users {
		fmt.Println("User:", name)
		for index, hobby := range hobbies {
			fmt.Printf("%v - Hobbie: %v\n", index, hobby)
		}
	}
}
