/*
	- Crie um valor e atribua-o a uma variável
	- Demonstre o endereço deste valor na memória
*/

package main

import "fmt"

func main() {
	example := "example"

	fmt.Println("Valor:", example)
	fmt.Println("Memória:", &example)
}
