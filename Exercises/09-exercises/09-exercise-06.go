/*
	- Crie um programa que demonstra seu OS e ARCH.
	- Rode com os seguitnes comandos:
		- go run
		- go build
		- go install
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é:", runtime.GOOS)
	fmt.Println("Seu OS é:", runtime.GOARCH)
}
