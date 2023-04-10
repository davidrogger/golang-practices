// Format print:
// - Decimal %d
// - Hexadecimal %#x
// - Unicode %#U
// - Tab \t
// - Linha nova \n
// Faça um loop dos números 33 a 122, e utilize format print para demonstrá-los como texto/string.

package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("- Decimal: %d \t - Hexadecimal: %#x \t - Unicode: %c\n", i, i, i)
	}
}
