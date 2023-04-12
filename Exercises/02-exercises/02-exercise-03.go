package main

import "fmt"

func main() {
	const a int = 10
	const b = 10

	fmt.Printf("%v %T", a, a)
	fmt.Printf("%v %T", b, b)
}
