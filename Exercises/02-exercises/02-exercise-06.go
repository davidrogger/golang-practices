package main

import (
	"fmt"
)

func main() {
	const (
		_ = 2023 + iota
		one
		two
		three
		four
	)

	fmt.Println(one, two, three, four)
}
