/*
	- [] Crie um tipo "quadrado"
	- [] Crie um tipo "circulo"
	- [] Crie um método "area" para tipo que calcule e retorne a área da figura
		- Área do círculo: 2* pi * raio
		- Área do quadrado: lado * lado
	- [] Crie um tipo "figura" que defina como interface qualquer tipo que
	tiver o método "área"
	- [] Crie uma função "info" que receba um tipo "figura" e retorne a área
	da figura
	- [] Crie um valor de tipo quadrado
	- [] Crie um valor de tipo "círculo"
	- Use a função "info" para demonstrar a área do "quadrado"
	- Use a função "info" para demonstrar a área do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	side float64
}

type circulo struct {
	radius float64
}

func (q quadrado) area() float64 {
	return q.side * q.side
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.radius
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{10}
	c := circulo{10}

	fmt.Println(info(q))
	fmt.Println(info(c))
}
