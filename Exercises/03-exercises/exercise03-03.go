// Crie um loop utilizando a sintaxe for condition
// Utilize para demonstrar os anos desde que você nasceu até uma data posterior
// Exemplo
// nasci no ano de 1988, até o ano de 2100

package main

import "fmt"

func main() {
	birthYear := 1988
	targetYear := 2100

	for birthYear <= targetYear {
		fmt.Println(birthYear)
		birthYear++
	}
}
