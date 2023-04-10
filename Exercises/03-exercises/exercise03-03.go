// Crie um loop utilizando a sintaxe for condition
// Utilize para demonstrar os anos desde que você nasceu

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
