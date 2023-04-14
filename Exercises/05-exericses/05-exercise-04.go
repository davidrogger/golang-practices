/**
	- Crie e use um struct anônimo
	- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
**/

package main

import "fmt"

func main() {
	jonas := struct {
		name   string
		skills []string
		grades map[string]int
	}{"jonas", []string{"sleep", "mock"}, map[string]int{"sleep": 10}}

	fmt.Println(jonas)
}
