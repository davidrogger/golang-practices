package main

import "fmt"

func main() {
	a := 10 == 1
	b := 10 != 1
	c := 10 <= 1
	d := 10 < 1
	e := 10 >= 1
	f := 10 > 1

	fmt.Println(a, b, c, d, e, f)
}
