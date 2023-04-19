/*
	- Atribua uma função a uma variável
	- Chame essa função
*/

package main

import "fmt"

func main() {
	x := func(x int) {
		fmt.Println(x)
	}

	x(10)
}
