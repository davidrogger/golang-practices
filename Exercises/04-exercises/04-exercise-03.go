/**
	- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
		- Do primeiro ao terceiro item do slice (incluindo o terceiro)
		- Quinto ao último item
		- Segundo ao sétimo
		- Terceiro ao penúltimo
		- Desafio obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
**/

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])

}
