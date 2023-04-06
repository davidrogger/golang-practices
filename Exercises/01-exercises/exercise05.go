package main

import "fmt"

type customInt int

var x customInt
var y int

func main() {
	fmt.Printf("ZeroValue: %v Type: %T\n", x, x)
	fmt.Printf("ZeroValue: %v Type: %T\n", y, y)
	x = 42
	y = int(x)

	fmt.Println(x)
	fmt.Println(y)
}
