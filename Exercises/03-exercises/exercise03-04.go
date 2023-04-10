// Crie um loop utilizando a sintaxe for
// Utilize para demonstrar os anos desde que você nasceu
// Exemplo
// nasci no ano de 1988, até o ano de 2100

package main

import "fmt"

func main() {
	birthYear := 1988
	targetYear := 2049

	for {
		if birthYear > targetYear {
			break
		} else {
			fmt.Println(birthYear)
			birthYear++
		}
	}
}
