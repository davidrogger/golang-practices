/*
	- Utilizando este código: https://go.dev/play/p/wlEM1tgfQD
	- ... use o struct sqrt.Erro como valor do tipo erro.
*/

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("Square root of negative number: %v", f)
		return 0, sqrtError{"00N", "00W", err}
	}
	return 42, nil
}

func main() {
	_, err := sqrt(-10.23)

	if err != nil {
		log.Println(err)
	}
}
