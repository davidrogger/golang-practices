package main

import "fmt"

type customInt int

var x customInt

func main() {
	fmt.Printf("ZeroValue: %v Type: %T\n", x, x)
	x = 42

	fmt.Println(x)
}
